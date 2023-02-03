//go:build wireinject

package di

import (
	"github.com/google/wire"
	"gitlab.kaonavi.jp/ae/kgm/auth"
	"gitlab.kaonavi.jp/ae/sardine/internal/cmd/api/middleware"
)

func InitializeAuthenticateToken() *middleware.AuthenticateToken {
	wire.Build(
		middleware.NewAuthenticateToken,
		auth.NewHS256Authenticator,
	)
	return nil
}
