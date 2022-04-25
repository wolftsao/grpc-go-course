package main

import (
	"context"
	"fmt"
	"log"

	"github.com/wolftsao/grpc-go-course/blog/blogpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("Blog Client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("couldn't not connect: %v", err)
	}

	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)

	fmt.Println("Creating the blog")
	blog := &blogpb.Blog{
		AuthorId: "Nathan",
		Title:    "My First Blog",
		Content:  "Content of the first blog",
	}

	ctx := context.Background()
	createBlogRes, err := c.CreateBlog(ctx, &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	fmt.Printf("Blog has been created: %v\n", createBlogRes)
	blogID := createBlogRes.GetBlog().GetId()

	fmt.Println("Reading the blog")
	_, err = c.ReadBlog(ctx, &blogpb.ReadBlogRequest{
		BlogId: "6266a97dc7s54ff84a6c247a",
	})
	if err != nil {
		fmt.Printf("Error happened while reading: %v\n", err)
	}

	readBlogReq := &blogpb.ReadBlogRequest{
		BlogId: blogID,
	}
	readBlogRes, err := c.ReadBlog(ctx, readBlogReq)
	if err != nil {
		fmt.Printf("Error happened while reading: %v\n", err)
	}

	fmt.Printf("Blog was read: %v\n", readBlogRes)

	newBlog := &blogpb.Blog{
		Id:       blogID,
		AuthorId: "Changed Author",
		Title:    "My First Blog (edited)",
		Content:  "Content of the first blog, with some awesome additions!",
	}

	updateRes, err := c.UpdateBlog(ctx, &blogpb.UpdateBlogRequest{
		Blog: newBlog,
	})
	if err != nil {
		fmt.Printf("Error happened while updating: %v\n", err)
	}

	fmt.Printf("Blog was read: %v\n", updateRes)

	deleteRes, err := c.DeleteBlog(ctx, &blogpb.DeleteBlogRequest{
		BlogId: blogID,
	})
	if err != nil {
		fmt.Printf("Error happened while deleting: %v\n", err)
	}
	fmt.Printf("Blog was deleted: %v\n", deleteRes)
}
