package accesstoken

import (
	"github.com/oauth/api/src/utils"
)

type Repository interface {
	GetById(string) (*AccessToken, *utils.RestErr)
}

type Service interface {
	GetById(string) (*AccessToken, *utils.RestErr)
}
type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}
func (s *service) GetById(accessTokenId string) (*AccessToken, *utils.RestErr) {
	accessToken, err := s.repository.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}
