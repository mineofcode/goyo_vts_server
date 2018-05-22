package db

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/go-playground/log"
	"goyo.in/gpstracker/client"
	opts "goyo.in/gpstracker/const"
)

var (
	Pool *client.Pool
)

type Factory struct {
	tile38Pool *client.Pool
}

func init() {
	// redisHost := os.Getenv("REDIS_HOST")
	// if redisHost == "" {
	// 	redisHost = opts.DefaultOpts.Tile38.Addr
	// }
	t38, err := initTile38Pool(opts.DefaultOpts.Tile38.Addr)
	if err != nil {
		log.WithFields(log.F("func", "NewFactory.initTile38Pool.Tile38")).Warn(err)
	}
	Pool = t38
	// cleanupHook()
}

func (f *Factory) Destory() {
	if f.tile38Pool != nil {
		err := f.tile38Pool.Close()
		if err != nil {
			log.WithFields(log.F("func", "Factory.Destory")).Warn(err)
		}
	}
}

//建立tile38连接池
func initTile38Pool(addr string) (*client.Pool, error) {
	p, err := client.DialPool(addr)
	if err != nil {
		log.WithFields(log.F("func", "initTile38Pool")).Warn(err)
		return nil, err
	}

	return p, nil
}

func newPool(server string) *redis.Pool {

	return &redis.Pool{

		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func cleanupHook() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGKILL)
	go func() {
		<-c
		Pool.Close()
		os.Exit(0)
	}()
}
