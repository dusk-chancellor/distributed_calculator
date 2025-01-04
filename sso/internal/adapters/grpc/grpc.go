package grpc

import (

	"github.com/dusk-chancellor/distributed_calculator/sso/internal/service"

	pb "github.com/dusk-chancellor/distributed_calculator/protos/gen/go/sso"
)

type Service interface {
	Register() ()
	Login() ()
	Logout() ()

	GetUser() ()
	UpdateUser() ()
	UpdateRole() ()
	ChangePassword() ()

	ValidateToken() ()
	RefreshToken() ()
}

type server struct {
	pb.UnimplementedSSOServiceServer
	service service.Service
}
