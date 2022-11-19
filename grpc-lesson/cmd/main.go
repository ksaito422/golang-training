package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	pb "grpc-lesson/pkg/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type myServer struct {
	pb.BackendServiceServer
}

func NewMyServer() *myServer {
	return &myServer{}
}

func main() {
	port := 8000
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterBackendServiceServer(s, NewMyServer())

	reflection.Register(s)

	go func() {
		log.Printf("start gRPC server port: %v", port)
		s.Serve(listener)
	}()

	// Ctrl+Cが入力されたらGraceful shutdownされるようにする
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}

func (s *myServer) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	// oneofの型判定
	switch v := req.Name.(type) {
	case *pb.HelloRequest_FullName:
		fmt.Println(v.FullName)
		call(ctx, v)
	case *pb.HelloRequest_LastName:
		fmt.Println(v.LastName)
	case *pb.HelloRequest_FirstName:
		fmt.Println(v.FirstName)
	default:
	}

	return &pb.HelloResponse{}, nil
}

// oneof switchの中で呼んでも、また型判定しないといけない
// コンパイル必要だからそうか
func call(ctx context.Context, v any) {
	switch i := v.(type) {
	case *pb.HelloRequest_FullName:
		fmt.Println(i.FullName)
	}
}
