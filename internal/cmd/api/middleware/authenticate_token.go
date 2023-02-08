package middleware

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	ka "gitlab.kaonavi.jp/ae/kgm/auth"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/authed"
	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/env"
	"gitlab.kaonavi.jp/ae/sardine/internal/ctxt"
	"gitlab.kaonavi.jp/ae/sardine/internal/errs"
	"gitlab.kaonavi.jp/ae/sardine/internal/utils/logger"
)

type AuthenticateToken struct {
	auth ka.Authenticator
}

func NewAuthenticateToken(a ka.Authenticator) *AuthenticateToken {
	return &AuthenticateToken{auth: a}
}

func (m *AuthenticateToken) Handler(ctx *gin.Context) {
	keys, err := env.GetTokenKey()
	if err != nil {
		logger.Error(ctx, err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": []string{"Internal Server Error"}})
		ctx.Abort()
		return
	}

	token, err := m.auth.ValidRequest(ctx.Request, ka.Key{
		Encrypt: keys.EncryptKey,
		Signing: keys.SigningKey,
	})
	if err != nil {
		if ka.IsValidationError(err) {
			ctx.JSON(http.StatusBadRequest, gin.H{"errors": []string{err.Error()}})
			ctx.Abort()
			return
		}

		logger.Error(ctx, errs.NewInternalError("トークン認証でエラーが発生しました: %v", err))
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": []string{"invalid token"}})
		ctx.Abort()
		return
	}

	authenticate := authed.New(
		token.Raw, // Raw = アクセストークン
		token.Payload.CustomerCode,
		token.Payload.LoginAddress,
		token.Payload.Client,
		token.Payload.LoginUserId,
		token.Payload.RoleType,
	)

	ctx.Request = ctx.Request.WithContext(
		context.WithValue(ctx.Request.Context(), ctxt.KeyAuthenticatedUser, &authenticate))
	ctx.Next()
}
