package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	pb "stargazers_grpc/stargazers"

	"google.golang.org/grpc"
)

const (
	port = ":8000"
)

type server struct {
	pb.UnimplementedStargazersServer
}

// FindStargazerCount performs an HTTP GET request against GitHub, parses the response
// and extracts the number of Stars against a given repository.
func FindStargazerCount(project, repo string) interface{} {
	githubEndpoint := "https://api.github.com/repos"
	repoEndpoint := fmt.Sprintf("%v/%v/%v", githubEndpoint, project, repo)
	response, err := http.Get(repoEndpoint)

	if err != nil {
		return fmt.Errorf("The HTTP request failed with error %s", err)
	}

	return ParseJSONForStargazers(response)
}

// ParseJSONForStargazers parses a json response returned from the GitHub API and returns the
// Stargazers count of a given repository
func ParseJSONForStargazers(httpResponse *http.Response) int64 {
	data, _ := ioutil.ReadAll(httpResponse.Body)

	var result map[string]interface{}
	json.Unmarshal([]byte(data), &result)

	starCount := result["stargazers_count"].(int64)
	return starCount
}

func (s *server) StargazersEndpoint(ctx context.Context, in *pb.StargazerRequest) (*pb.StargazersResponse, error) {
	stargazerCount := FindStargazerCount(in.GetProject(), in.GetRepository())
	return &pb.StargazersResponse{Count: stargazerCount.(int64)}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterStargazersServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
