package orchestrator

import (
	"context"
	"log"

	pb "github.com/dusk-chancellor/distributed_calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func Calculate(ctx context.Context, expr string, addr string) (float64, error) {

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("could not connect to grpc server:", err)
		return 0, status.Errorf(status.Code(err), "could not connect to grpc server: %v", err)
	}
	defer conn.Close()

	stat := conn.GetState().String()
	log.Println("connection state:", stat)

	grpcClient := pb.NewCalculatorServiceClient(conn)

	ans1, err := grpcClient.Calculate(ctx, &pb.ExpressionRequest{Expression: expr})
	if err != nil {
		log.Println("could not calculate:", err)
		return 0, status.Errorf(status.Code(err), "could not calculate: %v", err)
	}

	log.Printf("Agent response -> %v", ans1.Result)
	return ans1.Result, nil
}
