package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	pb "statement6/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("server-service:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewStringServiceClient(conn)
	ctx := context.Background()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter your chouce :Add/Remove/Check/List/Quit: ")
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimSpace(strings.ToLower(cmd))

		var input string
		if cmd != "list" && cmd != "quit" {
			fmt.Print("Input: ")
			input, _ = reader.ReadString('\n')
			input = strings.TrimSpace(input)
		}

		switch cmd {
		case "add":
			resp, _ := client.Add(ctx, &pb.StringRequest{Value: input})
			fmt.Println(resp.Message)
		case "remove":
			resp, _ := client.Remove(ctx, &pb.StringRequest{Value: input})
			fmt.Println(resp.Message)
		case "check":
			resp, _ := client.Check(ctx, &pb.StringRequest{Value: input})
			fmt.Println(resp.Message)
		case "list":
			resp, _ := client.List(ctx, &pb.Empty{})
			fmt.Println("Words:", resp.Values)
		case "quit":
			fmt.Println("Exiting client")
			return
		default:
			fmt.Println("Invalid command, TRY AGAIB!")
		}
	}
}
