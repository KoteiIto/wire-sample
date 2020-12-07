// AuthServiceをWireを使ってDIする場合
//+build wireinject

package di

import (
	"github.com/google/wire"

	"github.com/KoteiIto/wire-sample/pkg/repository"
	"github.com/KoteiIto/wire-sample/pkg/service"
)

func InitializeAuthService() *service.AuthService {
	wire.Build(
		service.NewAuthService,
		repository.NewOnetimeTokenRepository,
		repository.NewUserRepository,
		)
	return &service.AuthService{}
}