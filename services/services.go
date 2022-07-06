package services

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"

	"github.com/Muhammadjon226/api_gateway/config"
	pbPost "github.com/Muhammadjon226/api_gateway/genproto/post_service"
	pbFirst "github.com/Muhammadjon226/api_gateway/genproto/first_service"
)

type IServiceManager interface {
	PostService() pbPost.PostServiceClient
	FirstService() pbFirst.FirstServiceClient
}

type serviceManager struct {
	postService pbPost.PostServiceClient
	firstService pbFirst.FirstServiceClient
}

func (s *serviceManager) PostService() pbPost.PostServiceClient {
	return s.postService
}
func (s *serviceManager) FirstService() pbFirst.FirstServiceClient {
	return s.firstService
}

func NewServiceManager(conf *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	connPost, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.PostServiceHost, conf.PostServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	connFirst, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.FirstServiceHost, conf.FirstServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	serviceManager := &serviceManager{
		postService: pbPost.NewPostServiceClient(connPost),
		firstService: pbFirst.NewFirstServiceClient(connFirst),
	}

	return serviceManager, nil
}
