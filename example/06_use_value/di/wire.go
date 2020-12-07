// Provider providerを利用せずにDIをする
//+build wireinject

package di

import (
	"github.com/google/wire"

	"github.com/KoteiIto/wire-sample/pkg/repository"
	"github.com/KoteiIto/wire-sample/pkg/service"
)

var userRepository *repository.UserRepository

func init() {
	userRepository = repository.NewUserRepository()
}

func InitializeAuthService() *service.AuthService {
	wire.Build(
		service.NewAuthService,
		repository.NewOnetimeTokenRepository,
		wire.Value(userRepository),
	)
	return &service.AuthService{}
}
