package repo

import (
	"database/sql"

	"github.com/dusk-chancellor/distributed_calculator/sso/internal/models"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

type Repo struct {
	logger *zap.Logger
	db *sql.DB
}

func New(db *sql.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) CreateUser(username, email, role string, password []byte) (string, error) {
	return "", nil
}

func (r *Repo) GetUserByEmail(email string) (models.User, error) {
	return models.User{}, nil
}

func (r *Repo) GetUserByName(name string) (models.User, error) {
	return models.User{}, nil
}

func (r *Repo) GetUserById(id string) (models.User, error) {
	return models.User{}, nil
}

func (r *Repo) UpdateUser(models.User) (bool, error) {
	return false, nil
}
