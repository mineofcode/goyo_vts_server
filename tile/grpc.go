package tile

import (
	"fmt"
	"net"
	"sync/atomic"

	"github.com/go-playground/log"
	pb "goyo.in/gpstracker/hservice"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct {
	sCnt int64
}

func (s *server) Send(ctx context.Context, in *pb.MessageRequest) (*pb.MessageReply, error) {
	atomic.AddInt64(&s.sCnt, 1)
	// fmt.Println(in)
	// data := datamodel.GeofenceDetect{}
	// res := []datamodel.FenceTime{} // create varible

	// json.Unmarshal([]byte(in.Value), &data)
	// strconvs := data.Meta["time"].(string)

	// fmt.Println(strconvs)
	// err := json.Unmarshal([]byte(strconvs), &res)

	// if err == nil {
	// 	return &pb.MessageReply{Ok: true}, nil
	// }
	// fmt.Println(res)

	// go time.Sleep(5 * time.Second)
	//lg.Printf("Receive message %s", data.Detect)
	//go
	// redigogeofence.CallService(data)
	//log.WithFields(log.F("func", "server.Send"), log.F("sCnt", s.sCnt)).Info(in.String())
	return &pb.MessageReply{Ok: true}, nil
}

func GRpcRun() {

	fmt.Println("Start GRPC Server: ")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 6989))
	if err != nil {
		log.WithFields(log.F("func", "gRpcRun")).Fatal(err.Error())
	}
	s := grpc.NewServer()
	pb.RegisterHookServiceServer(s, &server{sCnt: 0})

	s.Serve(lis)
}
