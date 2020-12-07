// Provider ProviderがInterfaceを引数に取る場合は、Bindを利用する
//+build wireinject

package di

import (
	"github.com/google/wire"

	"github.com/KoteiIto/wire-sample/pkg/repository"
	"github.com/KoteiIto/wire-sample/pkg/service"
)

func InitializeAuthService() *service.AuthService {
	wire.Build(
		service.NewAuthServiceFromInterface,
		repository.NewUserRepository,
		wire.Bind(new(repository.IUserRepository), new(*repository.UserRepository)),
		repository.NewOnetimeTokenRepository,
		wire.Bind(new(repository.IOnetimeTokenRepository), new(*repository.OnetimeTokenRepository)),
	)
	return &service.AuthService{}
}
