package database

import (
	"context"
	"fmt"

	"gitlab.kaonavi.jp/ae/sardine/internal/core/infrastructure/env"
	"gitlab.kaonavi.jp/ae/sardine/internal/errs"
	"gitlab.kaonavi.jp/ae/sardine/internal/utils/logger"
	"gitlab.kaonavi.jp/ae/sardine/internal/utils/timer"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Conn はDB接続設定を保持するための構造体です
type Conn struct {
	// 接続先の顧客コード
	customerCode string
	// 接続設定（読み込み
	readSetting *env.DbConnectSetting
	// 接続設定（書き込み
	writeSetting *env.DbConnectSetting

	// 顧客ごとのReadコネクション
	dbRead *gorm.DB
}

func (c *Conn) init() error {
	db, err := c.open(c.readSetting)
	if err != nil {
		return errs.Wrap("[Conn.init]読み込みホストへの接続エラー", err)
	}
	c.dbRead = db
	return nil
}

// DB は読み込み用の接続設定を返却します
func (c *Conn) DB() *gorm.DB {
	return c.dbRead
}

// Transaction は書き込み処理を行います
// 書き込み用の接続情報を返却するinterfaceはないため、書き込みは必ずこのメソッドを通して行う必要があります
// クロージャに実際の書き込み処理を実装してください
func (c *Conn) Transaction(ctx context.Context, f func(tx *gorm.DB) error) (err error) {
	db, err := c.open(c.writeSetting)
	if err != nil {
		return errs.Wrap("[Conn.Transaction]読み込みホストへの接続エラー", err)
	}
	defer c.close(db)

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			logger.Panic(ctx, r)
			err = errs.NewInternalError("[Conn.Transaction]トランザクション処理中にpanicが発生しました: %v", r)
		}

		if err != nil {
			tx.Rollback()
		}
	}()

	if err = f(tx); err != nil {
		return errs.Wrap("[Conn.Transaction]Closureの実行エラー", err)
	}
	return tx.Commit().Error
}

// Close はDB接続を全て切断します
func (c *Conn) Close() {
	if c.dbRead != nil {
		c.close(c.dbRead)
	}
}

func (c *Conn) open(setting *env.DbConnectSetting) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=true&loc=%s",
		setting.Username,
		setting.Password,
		setting.Host,
		setting.Port,
		c.customerCode, // 接続先のDB名は顧客コード
		"Asia%2FTokyo",
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NowFunc: timer.Now,
	})
	if err != nil {
		return nil, errs.Wrap("[Conn.open]DB接続エラー", err)
	}
	return db, nil
}

func (*Conn) close(db *gorm.DB) {
	db2, err := db.DB()
	if err != nil {
		return
	}
	db2.Close()
}
