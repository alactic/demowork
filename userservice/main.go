package main

import (
	 "fmt"
	"context"

	proto "github.com/alactic/demosample/userservice/proto/user"
	"github.com/micro/go-micro"
)


type server struct{}

func main() {
	srv := micro.NewService(
		micro.Name("user"),
	)

	srv.Init()

	proto.RegisterUserServiceHandler(srv.Server(), &server{})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}

}

func (s *server) UserDetails(ctx context.Context, request *proto.Request, response *proto.Response) (error) {
	id := request.GetId()
    fmt.Println("id user :: ", id)
	// response.Consignments = id

	return nil
}
