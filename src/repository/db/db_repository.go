package db

import (
	accesstoken "github.com/oauth/api/src/domain/access_token"
	"github.com/oauth/api/src/utils"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type dbRepository struct {
}

type DbRepository interface {
	GetById(string) (*accesstoken.AccessToken, *utils.RestErr)
}

func (r *dbRepository) GetById(id string) (*accesstoken.AccessToken, *utils.RestErr) {
	return nil, utils.NewInternalServerError("database connection is not imp yet")
}
