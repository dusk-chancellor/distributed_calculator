package grpc

import (
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
	pb.
	service Service
}
