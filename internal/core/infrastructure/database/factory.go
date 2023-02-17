package database

import (
	"context"

	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/env"
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

func (*connFactory) Create(ctx context.Context) (*Conn, error) {
	authedUser, err := ctxt.AuthenticatedUser(ctx)
	if err != nil {
		return nil, errs.Wrap("[connFactory.Create]ctxt.AuthenticatedUserのエラー", err)
	}

	readSetting, err := env.GetReadDbConnectSetting()
	if err != nil {
		return nil, errs.Wrap("[connFactory.Create]env.GetReadDbConnectSettingのエラー", err)
	}
	writeSetting, err := env.GetWriteDbConnectSetting()
	if err != nil {
		return nil, errs.Wrap("[connFactory.Create]env.GetWriteDbConnectSettingのエラー", err)
	}

	conn := &Conn{
		customerCode: authedUser.CustomerCode(),
		readSetting:  readSetting,
		writeSetting: writeSetting,
	}
	if err = conn.init(); err != nil {
		return nil, errs.Wrap("[connFactory.Create]Conn.initのエラー", err)
	}
	return conn, nil
}
