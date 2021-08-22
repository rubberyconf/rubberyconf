package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/rubberyconf/rubberyconf/grpcapi/grpcapipb"
	"github.com/rubberyconf/rubberyconf/grpcapi/servers"
	"google.golang.org/grpc"
)

func loadConfiguration(path string) *config.Config {

	environment := os.Getenv("ENV")
	if environment == "" {
		environment = "local"
	}

	conf := config.NewConfiguration(filepath.Join(path, fmt.Sprintf("../../config/%s.yml", environment)))
	b, _ := json.MarshalIndent(conf, "", "   ")
	logs.GetLogs().WriteMessage("debug", fmt.Sprintf("Configuration loaded:\n%s\nEnvironment: %s ", string(b), environment), nil)
	return conf
}
func main() {

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
		panic(err)
	}
	conf := loadConfiguration(path)

	address := "0.0.0.0:50051"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Server is listening on %v ...", address)

	s := grpc.NewServer()
	grpcapipb.RegisterRubberyConfServiceServer(s, &servers.ConfServer{})
	grpcapipb.RegisterRubberyFeatureServiceServer(s, &servers.FeatureServer{})

	s.Serve(lis)
}
