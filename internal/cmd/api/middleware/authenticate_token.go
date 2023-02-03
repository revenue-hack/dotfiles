package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	ka "gitlab.kaonavi.jp/ae/kgm/auth"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/model/auth"
	"gitlab.kaonavi.jp/ae/sardine/internal/ctxt"
	"gitlab.kaonavi.jp/ae/sardine/internal/utils/logger"
)

type AuthenticateToken struct {
	auth ka.Authenticator
}

func NewAuthenticateToken(a ka.Authenticator) *AuthenticateToken {
	return &AuthenticateToken{auth: a}
}

func (m *AuthenticateToken) Handler(ctx *gin.Context) {
	token, err := m.auth.ValidRequest(ctx.Request, ka.Key{
		Encrypt: os.Getenv("TOKEN_ENCRYPT_KEY"),
		Signing: os.Getenv("TOKEN_SIGNING_KEY"),
	})

	ctx.Request.Context()

	if err != nil {
		if ka.IsValidationError(err) {
			ctx.JSON(http.StatusBadRequest, gin.H{"errors": []string{err.Error()}})
			ctx.Abort()
			return
		}

		logger.Error(ctx, fmt.Errorf("トークン認証でエラーが発生しました: %s", err.Error()))
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": []string{"invalid token"}})
		ctx.Abort()
		return
	}

	authenticate := auth.New(
		token.Raw, // Raw = アクセストークン
		token.Payload.CustomerCode,
		token.Payload.LoginAddress,
		token.Payload.Client,
		token.Payload.LoginUserId,
		token.Payload.RoleType,
	)

	ctx.Request = ctx.Request.WithContext(context.WithValue(ctx.Request.Context(), ctxt.KeyAuthenticatedUser, authenticate))
	ctx.Next()
}
