// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package di

import (
	"github.com/KoteiIto/wire-sample/pkg/repository"
	"github.com/KoteiIto/wire-sample/pkg/service"
)

// Injectors from wire.go:

func InitializeAuthService() (*service.AuthService, func()) {
	userRepository := repository.NewUserRepository()
	onetimeTokenRepository := repository.NewOnetimeTokenRepository()
	authService, cleanup := service.NewAuthServiceWithCleanup(userRepository, onetimeTokenRepository)
	return authService, func() {
		cleanup()
	}
}