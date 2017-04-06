package main

import (
  "flag"
  "net/http"

  "github.com/golang/glog"
  "github.com/golang/net/context"
  // "golang.org/x/net/context"
  "github.com/grpc-ecosystem/grpc-gateway/runtime"
  // "github.com/grpc/grpc"
  "google.golang.org/grpc"

  // gw "github.com/SolarisYan/grpc-hello-with-gateway"
  gw "helloworld"
)

var (
  echoEndpoint = flag.String("helloworld_endpoint", "localhost:50051", "endpoint of Greeter gRPC Service")
)

func run() error {
  ctx := context.Background()
  ctx, cancel := context.WithCancel(ctx)
  defer cancel()

  mux := runtime.NewServeMux()
  opts := []grpc.DialOption{grpc.WithInsecure()}
  err := gw.RegisterGreeterHandlerFromEndpoint(ctx, mux, *echoEndpoint, opts)
  if err != nil {
    return err
  }

    // glog.Print("Greeter gRPC Server gateway start at port 8080...")
    return http.ListenAndServe(":8080", mux)
}

func main() {
  flag.Parse()
  defer glog.Flush()

  if err := run(); err != nil {
    glog.Fatal(err)
  }
}