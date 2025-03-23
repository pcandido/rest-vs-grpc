package main

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"net/http"

	"github.com/pcandido/rest-vs-grpc/proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedComplexDataServiceServer
}

func main() {
	go restServer()
	go grpcServer()

	select {}
}

func restServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := GenerateComplexData()
		jsonData, _ := json.MarshalIndent(data, "", "  ")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	})

	log.Println("Starting REST server at port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func grpcServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("falha ao escutar: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterComplexDataServiceServer(s, &server{})
	reflection.Register(s)

	log.Println("Starting gRPC server at port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("falha ao iniciar servidor gRPC: %v", err)
	}
}

func (s *server) GetData(ctx context.Context, req *pb.ComplexDataRequest) (*pb.ComplexDataResponse, error) {
	data := GenerateComplexData()
	return data, nil
}
