package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"gogrpc/blogpb"
	"google.golang.org/grpc"
)

func ping() {

	serverAddress := "localhost:50051"

	conn, e := grpc.Dial(serverAddress, grpc.WithInsecure())
	if e != nil {
		panic(e)
	}
	defer conn.Close()

	client := blogpb.NewBlogServiceClient(conn)

	list, err := client.ListBlogs(context.Background(), &blogpb.ListBlogsReq{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for {
		item, err := list.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Get blogs failed")
		}
		log.Printf("Blog: %v", item)
	}
}

func main() {
	ping()
}

