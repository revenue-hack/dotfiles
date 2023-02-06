package database

import (
	"context"

	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/env"
	"gitlab.kaonavi.jp/ae/sardine/internal/ctxt"
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
		return nil, err
	}

	readSetting, err := env.GetReadDbConnectSetting()
	if err != nil {
		return nil, err
	}
	writeSetting, err := env.GetWriteDbConnectSetting()
	if err != nil {
		return nil, err
	}

	return &Conn{
		customerCode: authedUser.CustomerCode(),
		readSetting:  readSetting,
		writeSetting: writeSetting,
	}, nil
}
