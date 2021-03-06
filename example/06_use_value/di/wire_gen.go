// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package di

import (
	"github.com/KoteiIto/wire-sample/pkg/repository"
	"github.com/KoteiIto/wire-sample/pkg/service"
)

// Injectors from wire.go:

func InitializeAuthService() *service.AuthService {
	repositoryUserRepository := _wireUserRepositoryValue
	onetimeTokenRepository := repository.NewOnetimeTokenRepository()
	authService := service.NewAuthService(repositoryUserRepository, onetimeTokenRepository)
	return authService
}

var (
	_wireUserRepositoryValue = userRepository
)

// wire.go:

var userRepository *repository.UserRepository

func init() {
	userRepository = repository.NewUserRepository()
}
