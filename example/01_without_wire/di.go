// AuthServiceをWireを使わずに手動でDIする場合

package main

import (
	"github.com/KoteiIto/wire-sample/pkg/repository"
	"github.com/KoteiIto/wire-sample/pkg/service"
)

func InitializeAuthService() *service.AuthService {
	userRepository := repository.NewUserRepository()
	onetimeTokenRepository := repository.NewOnetimeTokenRepository()
	authService := service.NewAuthService(userRepository, onetimeTokenRepository)
	return authService
}
