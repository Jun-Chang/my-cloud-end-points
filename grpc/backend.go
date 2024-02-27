package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"connectrpc.com/connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	greetv1 "github.com/Jun-Chang/my-cloud-endpoints/grpc/gen/greet/v1"
	"github.com/Jun-Chang/my-cloud-endpoints/grpc/gen/greet/v1/greetv1connect"
)

type GreetServer struct{}

func (s *GreetServer) Greet(
	ctx context.Context,
	req *connect.Request[greetv1.GreetRequest],
) (*connect.Response[greetv1.GreetResponse], error) {
	log.Println("Request headers:", req.Header())
	res := connect.NewResponse(&greetv1.GreetResponse{
		Greeting: fmt.Sprint("Hello, ", req.Msg.Name),
	})
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}

func main() {
	greeter := &GreetServer{}
	mux := http.NewServeMux()
	path, hander := greetv1connect.NewGreetServiceHandler(greeter)
	mux.Handle(path, hander)
	http.ListenAndServe(
		"0.0.0.0:8080",
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
