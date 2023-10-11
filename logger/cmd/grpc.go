package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/mstgnz/microservice/proto"
	"google.golang.org/grpc"
)

type LogServer struct {
	proto.UnimplementedLogServiceServer
	Models Models
}

func (l *LogServer) WriteLog(ctx context.Context, req *proto.LogRequest) (*proto.LogResponse, error) {

	// write the log
	logEntry := LogEntry{
		Name: req.GetName(),
		Data: req.GetData(),
	}

	err := l.Models.LogEntry.Insert(logEntry)
	if err != nil {
		res := &proto.LogResponse{Result: false}
		return res, err
	}

	// return response
	res := &proto.LogResponse{Result: true}
	return res, nil
}

func (app *Config) gRPCListen() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", 50001))
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}

	s := grpc.NewServer()

	proto.RegisterLogServiceServer(s, &LogServer{Models: app.Models})

	log.Printf("gRPC Server started on port %v", 50001)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}
}
