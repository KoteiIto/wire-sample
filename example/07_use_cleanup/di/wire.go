// AuthServiceをWireを使ってDIする場合
//+build wireinject

package di

import (
	"github.com/google/wire"

	"github.com/KoteiIto/wire-sample/pkg/repository"
	"github.com/KoteiIto/wire-sample/pkg/service"
)

func InitializeAuthService() (*service.AuthService, func()) {
	wire.Build(
		service.NewAuthServiceWithCleanup,
		repository.NewOnetimeTokenRepository,
		repository.NewUserRepository,
	)
	return &service.AuthService{}, func() {}
}
