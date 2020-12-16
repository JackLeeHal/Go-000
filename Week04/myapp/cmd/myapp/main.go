package main

import (
	"context"
	"fmt"
	"log"
	pb "myapp/api/myapp"
	"net"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp",
		fmt.Sprintf("localhost:%d", 9999))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	client, err := InitClient(context.TODO(), os.Getenv("DB_HOST"))
	if err != nil {
		log.Fatal(err)
	}

	handle := InitProtoServe(client.Database("demo").Collection("demo"))
	pb.RegisterDemoServer(grpcServer, handle)

	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		fmt.Println("grpc")
		go func() {
			<-ctx.Done()
			fmt.Println("grpc ctx done")
			grpcServer.Stop()
		}()
		return grpcServer.Serve(lis)
	})

	// signal
	g.Go(func() error {
		exitSignals := []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT}
		sig := make(chan os.Signal, len(exitSignals))
		signal.Notify(sig, exitSignals...)
		for {
			fmt.Println("signal")
			select {
			case <-ctx.Done():
				fmt.Println("signal ctx done")
				return ctx.Err()
			case <-sig:
				// do something
				return nil
			}
		}
	})

	log.Fatal(g.Wait())
}
