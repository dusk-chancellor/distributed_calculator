package agent

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/dusk-chancellor/distributed_calculator/internal/utils/agent/calculation"
	itp "github.com/dusk-chancellor/distributed_calculator/internal/utils/agent/infix_to_postfix"
	pb "github.com/dusk-chancellor/distributed_calculator/proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.CalculatorServiceServer
}

func NewServer() *Server {
	return &Server{}
}

type CalculatorServiceServer interface {
	Calculate(context.Context, *pb.ExpressionRequest) (*pb.ExpressionResponse, error)
	mustEmbedUnimplementedCalculatorServiceServer()
}

func (s *Server) Calculate(ctx context.Context, in *pb.ExpressionRequest) (*pb.ExpressionResponse, error) {
	postfixed := itp.ToPostfix(in.Expression)
	res, err := calculation.Evaluate(postfixed)
	if err != nil {
		return nil, err
	}
	log.Println("successful operation!")
	return &pb.ExpressionResponse{Result: res}, nil
}

func RunAgentServer() {
	addr := fmt.Sprintf("%s:%s", "localhost", "5000")
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Println("error starting tcp listener: ", err)
		os.Exit(1)
	}

	log.Println("tcp listener started")

	grpcServer := grpc.NewServer()

	expressionServiceServer := NewServer()

	pb.RegisterCalculatorServiceServer(grpcServer, expressionServiceServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Println("error serving grpc: ", err)
		os.Exit(1)
	}
}