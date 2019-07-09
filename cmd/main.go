package main

import (
	"github.com/noguchi-hiroshi/grpc-gateway-test/client"
	gateway "github.com/noguchi-hiroshi/grpc-gateway-test/proto"
	"google.golang.org/grpc"
	"log"
)

func request(c client.UserClient) {
	id, err := c.Create("abc", "def")
	if err != nil {
		log.Fatalln(err)
	}
	entity, err := c.Find(id)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(entity)
}
func newGRPCClient() client.UserClient {
	conn, err := grpc.Dial(":4949", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial %v\n", err)
	}
	sc := gateway.NewUserServiceClient(conn)
	return client.NewGRPCUserClient(sc)
}

func newHTTPClient() client.UserClient {
	return client.NewHttpUserClient()
}

func main() {
	request(newGRPCClient())
	request(newHTTPClient())
}
