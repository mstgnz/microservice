package proto

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func WriteGRPCLog(name, data string) (*LogResponse, error) {
	dial, err := grpc.Dial("logger-service:50001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	defer func(dial *grpc.ClientConn) {
		_ = dial.Close()
	}(dial)

	d := NewLogServiceClient(dial)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	log, err := d.WriteLog(ctx, &LogRequest{Name: name, Data: data})

	return log, err

}
