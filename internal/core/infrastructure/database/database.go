package database

import (
	"context"
	"fmt"
	"os"

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
	// 顧客ごとのReadコネクション
	dbRead *gorm.DB
}

// Read は読み込み用の接続設定を返却します
func (c *Conn) Read() (*gorm.DB, error) {
	if c.dbRead != nil {
		return c.dbRead, nil
	}
	db, err := c.open(os.Getenv("DB_READ_HOST_NAME"))
	if err != nil {
		return nil, errs.NewInternalError("DBの読み込み用のホストへの接続に失敗しました: %v", err)
	}
	c.dbRead = db
	return c.dbRead, nil
}

// Transaction は書き込み処理を行います
// 書き込み用の接続情報を返却するinterfaceはないため、書き込みは必ずこのメソッドを通して行う必要があります
// クロージャに実際の書き込み処理を実装してください
func (c *Conn) Transaction(ctx context.Context, f func(tx *gorm.DB) error) (err error) {
	db, err := c.open(os.Getenv("DB_WRITE_HOST_NAME"))
	if err != nil {
		return errs.NewInternalError("DBの書き込み用のホストへの接続に失敗しました: %v", err)
	}
	defer c.close(db)

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			logger.Panic(ctx, r)
			err = errs.NewInternalError("トランザクション処理中にpanicが発生しました: %v", r)
		}

		if err != nil {
			tx.Rollback()
		}
	}()

	if err = f(tx); err != nil {
		return err
	}
	return tx.Commit().Error
}

// Close はDB接続を全て切断します
func (c *Conn) Close() {
	if c.dbRead != nil {
		c.close(c.dbRead)
	}
}

func (c *Conn) open(host string) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=true&loc=%s",
		os.Getenv("DB_USER_NAME"),
		os.Getenv("DB_PASSWORD"),
		host,
		os.Getenv("DB_PORT"),
		c.customerCode, // 接続先のDB名は顧客コード
		"Asia%2FTokyo",
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NowFunc: timer.Now,
	})
	if err != nil {
		return nil, errs.NewInternalError("failed to open db connection: %v", err)
	}
	return db, nil
}

func (c *Conn) close(db *gorm.DB) {
	db2, err := db.DB()
	if err != nil {
		return
	}
	db2.Close()
}
