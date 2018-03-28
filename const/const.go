package consts

import "time"

var opts *Options

const (
	BeegoMode = "prod" // dev || prod

	MGODbip   = "127.0.0.1"
	MGODbport = "27017"

	TCPHost   = "0.0.0.0"
	TCPPort   = 6969
	HTTPPort  = 6979
	HTTPSPort = 43
)

type (
	Options struct {
		Tile38  Tile38Opts
		Db      MongoDB
		WebScok WebSockets
	}

	Tile38Opts struct {
		Addr      string        `help:"tile38 server address(addr:port)" from:"env,flag"`
		MaxActive int           `help:"tile38 connection pool max active count" from:"env,flag"`
		IdleTime  time.Duration `help:"tile38 connection pool keepalive duration" from:"env,flag"`
	}

	WebSockets struct {
		Addr string `help:"Websocket server address(addr:port)" from:"env,flag"`
	}

	MongoDB struct {
		MGODbip   string `""`
		MGODbport string `""`
	}
)

var DefaultOpts = &Options{

	WebScok: WebSockets{
		Addr: ":6989",
	},
	Tile38: Tile38Opts{
		Addr:      ":9851",
		MaxActive: 10,
		IdleTime:  2 * time.Second,
	},
	Db: MongoDB{
		MGODbip:   "127.0.0.1",
		MGODbport: "27017",
	},
}
