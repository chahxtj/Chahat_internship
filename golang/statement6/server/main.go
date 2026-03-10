package main

import (
	"context"
	"log"
	"net"

	pb "statement6/proto"
	"statement6/trie"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedStringServiceServer
	trie *trie.Trie
}

func (s *server) Add(ctx context.Context, req *pb.StringRequest) (*pb.StringResponse, error) {
	s.trie.Add(req.Value)
	return &pb.StringResponse{Message: "Word Added!"}, nil
}

func (s *server) Remove(ctx context.Context, req *pb.StringRequest) (*pb.StringResponse, error) {
	s.trie.Delete(req.Value)
	return &pb.StringResponse{Message: "Word Removed!"}, nil
}

func (s *server) Check(ctx context.Context, req *pb.StringRequest) (*pb.StringResponse, error) {
	if s.trie.Exists(req.Value) {
		return &pb.StringResponse{Message: "WordExists!"}, nil
	}
	return &pb.StringResponse{Message: "Word Not Found"}, nil
}

func (s *server) List(ctx context.Context, req *pb.Empty) (*pb.ListResponse, error) {
	return &pb.ListResponse{Values: s.trie.List()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterStringServiceServer(grpcServer, &server{trie: trie.NewTrie()})

	log.Println("gRPC server is running on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
