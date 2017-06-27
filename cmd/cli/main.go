package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/agonzalezro/alex-gokit-example/pb"
	"google.golang.org/grpc"

	"flag"

	"io/ioutil"
)

func main() {
	httpAddr := flag.String("httpAddr", "http://localhost:8080/bye/alex", "the http address")
	grpcAddr := flag.String("grcpAddr", ":8081", "the gRPC address")
	flag.Parse()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		conn, err := grpc.Dial(*grpcAddr, grpc.WithInsecure(), grpc.WithTimeout(time.Second))
		if err != nil {
			log.Fatal("error dialing grpc: ", err)
		}
		defer conn.Close()

		client := pb.NewHelloClient(conn)
		v, err := client.Hi(context.TODO(), &pb.HiRequest{Name: "Alex"})
		if err != nil {
			log.Fatal("error calling Hi: ", err)
		}
		fmt.Println(v.V)
	}()

	go func() {
		defer wg.Done()

		resp, err := http.Get(*httpAddr)
		if err != nil {
			log.Fatalf("error getting %s: %v\n", *httpAddr, err)
		}
		defer resp.Body.Close()

		bs, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal("error reading the HTTP body: ", err)
		}
		fmt.Println(string(bs)) // We could unmarhsal, but meh
	}()

	wg.Wait()
}
