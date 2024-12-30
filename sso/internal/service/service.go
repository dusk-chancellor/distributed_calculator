package service

import (
	"github.com/dusk-chancellor/distributed_calculator/sso/internal/config"
	"github.com/dusk-chancellor/distributed_calculator/sso/internal/models"
	"go.uber.org/zap"
)

type Repo interface {
	CreateUser(username, email, role string, password []byte) (id string, err error)
	GetUserByEmail(email string) (models.User, error)
	GetUserByName(name string) (models.User, error)
	GetUserById(id string) (models.User, error)
	UpdateUser(models.User) (bool, error)
}

type Service struct {
	logger *zap.Logger
	user Repo
	jwt config.JWT
}

func New(logger *zap.Logger, user Repo, jwt config.JWT) *Service {
	return &Service{
		logger: logger,
		user:   user,
		jwt:    jwt,
	}
}


