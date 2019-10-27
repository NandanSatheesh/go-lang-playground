package main

import (
	"fmt"
	"github.com/NandanSatheesh/go-lang-playground/grpc/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"strconv"
)

func main() {

	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	clientAdd := proto.NewAddServiceClient(conn)
	clientMultiply := proto.NewMultiplyServiceClient(conn)

	g := gin.Default()
	g.GET("/add/:a/:b", func(ctx *gin.Context) {
		a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter A"})
			return
		}

		b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter B"})
			return
		}

		req := &proto.Request{A: int64(a), B: int64(b)}
		if response, err := clientAdd.Add(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(response.Result),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	g.GET("/mult/:a/:b", func(ctx *gin.Context) {
		a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter A"})
			return
		}
		b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter B"})
			return
		}
		req := &proto.Request{A: int64(a), B: int64(b)}

		if response, err := clientMultiply.Multiply(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(response.Result),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
