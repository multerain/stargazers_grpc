package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "stargazers_grpc/stargazers"

	"google.golang.org/grpc"
)

const (
	address = "localhost:8000"
)

var (
	project = flag.String("project", "EbookFoundation", "GitHub Project associated with repo to query, defaults to EbookFoundation.")
	repo    = flag.String("repo", "free-programming-books", "GitHub Repository to query Stargazer count, defaults to free-programming-books.")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewStargazersClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetStargazers(ctx, &pb.StargazerRequest{Project: *project, Repository: *repo})
	if err != nil {
		log.Fatalf("Could not get response: %v", err)
	}
	log.Printf("Stargazers: %v", r.GetCount())
}
