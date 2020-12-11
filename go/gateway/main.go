package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	gw "iminders/gateway/gen/go/proto/echo_service"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9090", "gRPC server endpoint")
)

type EchoServer struct {
	gw.UnimplementedEchoServiceServer
}

func (s *EchoServer) Echo(c context.Context, req *gw.StringMessage) (rep *gw.StringMessage, err error) {
	fmt.Println(req.Value)
	return &gw.StringMessage{Value: "hello " + req.Value}, nil
}

func (s *EchoServer) Downfile(req *gw.StringMessage, rep gw.EchoService_DownfileServer) (err error) {
	fmt.Println(req.Value)
	if req.Value == "" {
		return errors.New("input params empty")
	}
	fileBuffer, err := ioutil.ReadFile(req.Value)
	if err != nil {
		log.Println("err:", err.Error())
		return err
	}
	if err := rep.Send(&gw.FileBinary{
		Data: fileBuffer,
	}); err != nil {
		log.Println("send failed:", err.Error())
		return err
	}

	byteOnce := 50
	index := 0
	cnt := 0
	for {
		if index*byteOnce+byteOnce >= len(fileBuffer) {
			if err := rep.Send(&gw.FileBinary{
				Data: fileBuffer[index*byteOnce:],
			}); err != nil {
				log.Println("send failed:", err.Error())
				return err
			}
			cnt += len(fileBuffer) - index*byteOnce
			break
		} else {
			if err := rep.Send(&gw.FileBinary{
				Data: fileBuffer[index*byteOnce : index*byteOnce+byteOnce],
			}); err != nil {
				log.Println("send failed:", err.Error())
				return err
			}
			cnt += byteOnce
		}
		index++
	}
	log.Println("send bytes len:", cnt)
	return nil
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	lis, err := net.Listen("tcp", *grpcServerEndpoint)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer()
	gw.RegisterEchoServiceServer(s, &EchoServer{})
	go s.Serve(lis)

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err = gw.RegisterEchoServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":8081", mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
