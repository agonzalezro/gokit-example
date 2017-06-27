package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/agonzalezro/hiworld/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8081", grpc.WithInsecure(), grpc.WithTimeout(time.Second))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
		os.Exit(1)
	}
	defer conn.Close()

	client := pb.NewHelloClient(conn)
	v, err := client.Hi(context.TODO(), &pb.HiRequest{Name: "Alex"})
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(v.V)
}
