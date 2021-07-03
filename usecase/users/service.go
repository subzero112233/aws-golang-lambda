package users

import (
	"github.com/rs/zerolog"
)

type Service struct {
	Repository Repository
	Logger     *zerolog.Logger
}

func LoadService(repository Repository, logger *zerolog.Logger) *Service {
	return &Service{
		Logger:     logger,
		Repository: repository,
	}
}
