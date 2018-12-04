package client

import (
	"fmt"
	"net/url"
	"strings"

	pb "../protos"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	serverAddr = "127.0.0.1"
	grpcPort   = ":50051"
)

func SendRequest(articleUrl string) string {

	if (ValidateUrl(articleUrl)) == false {
		fmt.Println("Invalid Url Format")
		return ""
	}

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	conn, err := grpc.Dial(serverAddr+grpcPort, opts...)
	if err != nil {
		fmt.Printf("Failed to dial: %v\n", err)
		return ""
	}
	defer conn.Close()
	client := pb.NewParseArticleClient(conn)

	response, err := client.ParseUrl(context.Background(), &pb.ParseRequest{
		Url: articleUrl,
	})
	if err != nil {
		fmt.Printf("Failed to parse article: %v\n", err)
		return ""
	}

	var output = []string{"content : " + response.Article, "title : " + response.Title, "server_message : " + response.ServerMessage}
	stringSlices := strings.Join(output[:], "\n\n")

	return stringSlices

}

func ValidateUrl(urlToRead string) bool {
	_, err := url.ParseRequestURI(urlToRead)
	if err != nil {
		return false
	}
	return true
}
