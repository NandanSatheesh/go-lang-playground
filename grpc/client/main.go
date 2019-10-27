package main

import (
	"github.com/NandanSatheesh/go-lang-playground/grpc/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
)

func main() {

	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	clientAdd := proto.NewAddServiceClient(conn)
	clientMultiply := proto.NewMultiplyServiceClient(conn)

	g := gin.Default()

	g.GET("/add", func(i *gin.Context) {

		response, err := clientAdd.Add(i, &proto.Request{
			A: int64(10),
			B: int64(12),
		})

		if err != nil {
			panic(err)
		}

		print(response.Result)

	})

	g.GET("/mul", func(i *gin.Context) {

		response, err := clientMultiply.Multiply(i, &proto.Request{
			A: int64(10),
			B: int64(12),
		})

		if err != nil {
			panic(err)
		}

		print(response.Result)

	})

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
