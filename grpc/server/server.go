package main

import (
	"fmt"
	"net"
	"strings"

	"../../parser"
	"../protos"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	grpcPort = ":50051"
)

type Server struct{}

func (s *Server) ParseUrl(context context.Context, request *protos.ParseRequest) (*protos.ParseResponse, error) {
	title, content, message := parser.ParseFromUrl(request.Url)
	return &protos.ParseResponse{Title: title, Article: content, ServerMessage: strings.Join(message, ",")}, nil
}
func main() {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		fmt.Printf("failed to parse url: %v\n", err)
		return
	}
	grpcServer := grpc.NewServer()
	protos.RegisterParseArticleServer(grpcServer, &Server{})
	reflection.Register(grpcServer)
	grpcServer.Serve(listen)
	return
}
