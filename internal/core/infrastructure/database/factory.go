package database

import (
	"context"

	"gitlab.kaonavi.jp/ae/sardine/internal/ctxt"
	"gitlab.kaonavi.jp/ae/sardine/internal/errs"
)

type ConnFactory interface {
	Create(ctx context.Context) (*Conn, error)
}

func NewFactory() ConnFactory {
	return &connFactory{}
}

type connFactory struct{}

func (c *connFactory) Create(ctx context.Context) (*Conn, error) {
	authedUser, err := ctxt.AuthenticatedUser(ctx)
	if err != nil {
		return nil, errs.NewInternalError("failed to ctxt.AuthenticatedUser from connFactory: %v", err)
	}
	return &Conn{customerCode: authedUser.CustomerCode()}, nil
}
