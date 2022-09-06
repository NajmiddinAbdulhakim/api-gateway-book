package service

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"

	"github.com/NajmiddinAbdulhakim/ude/api-gateway/config"
	pb "github.com/NajmiddinAbdulhakim/ude/api-gateway/genproto"
)

type IServiceManager interface {
	BookService() pb.BookServcieClient
}

type serviceManager struct {
	bookService pb.BookServcieClient
}

func (s *serviceManager) BookService() pb.BookServcieClient {
	return s.bookService
}

func NewServiceManager(conf config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	conn, err := grpc.Dial(
		fmt.Sprintf(`%s:%d`, conf.BookServiceHost, conf.BookServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	serviceManager := &serviceManager{
		bookService: pb.NewBookServcieClient(conn),
	}
	return serviceManager, nil
}
