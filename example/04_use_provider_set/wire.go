// Provider Setを利用する
//+build wireinject

package main

import (
	"github.com/google/wire"

	"github.com/KoteiIto/wire-sample/pkg/repository"
	"github.com/KoteiIto/wire-sample/pkg/service"
)

func InitializeAuthService() *service.AuthService {
	wire.Build(
		service.NewAuthService,
		repository.RepositorySet,
	)
	return &service.AuthService{}
}
