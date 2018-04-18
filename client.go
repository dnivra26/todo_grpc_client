package main

import (
	"google.golang.org/grpc"
	"log"
	"todo_proto/pb/proto"
	"golang.org/x/net/context"
)

func main()  {
	conn, e := grpc.Dial(":7777", grpc.WithInsecure())
	if e != nil {
		log.Fatalf("did not connect: %s", e)
	}
	defer conn.Close()
	client := proto.NewTodoServiceClient(conn)
	response, err := client.GetTodo(context.Background(), &proto.GetTodoRequest{Id: "1"})
	if err != nil {
		log.Fatalf("Error when calling GetTodo: %s", err)
	}
	log.Printf("Response from server: %s", response.Todo)

}
