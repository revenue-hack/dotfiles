package database

import (
	"context"
	"errors"

	"gitlab.kaonavi.jp/ae/sardine/internal/errs"
	"gorm.io/gorm"
)

var (
	errRecordNotFound = errs.NewNotFound("record not found")
)

// IsErrRecordNotFound はレコードが存在しない場合のエラーである場合にtrueを返します
func IsErrRecordNotFound(err error) bool {
	return errors.Is(errRecordNotFound, err)
}

// Get は単一レコードの検索を行った結果を返します
func Get[T any](_ context.Context, db *gorm.DB) (*T, error) {
	var s T
	ret := db.Limit(1).Find(&s)
	if ret.Error != nil {
		return nil, errs.Wrap("[database.Get]データ取得エラー", ret.Error)
	}
	if ret.RowsAffected == 0 {
		return nil, errRecordNotFound
	}
	return &s, nil
}

// GetAll は複数レコードの検索を行った結果を返します
func GetAll[T any](_ context.Context, db *gorm.DB) ([]T, error) {
	var s []T
	if err := db.Find(&s).Error; err != nil {
		return nil, errs.Wrap("[database.GetAll]データ取得エラー", err)
	}
	return s, nil
}

// Exist は検索を行った結果が存在する場合はnil、存在しない場合はerrorを返却します
func Exist[T any](_ context.Context, db *gorm.DB) (bool, error) {
	var s T
	ret := db.Limit(1).Find(&s)
	if ret.Error != nil {
		return false, errs.Wrap("[database.Exist]データ取得エラー", ret.Error)
	}
	return ret.RowsAffected > 0, nil
}

// ExistById は指定IDのレコードが存在する場合はnil、存在しない場合はerrorを返却します
func ExistById[T any](ctx context.Context, db *gorm.DB, id uint32) (bool, error) {
	var s []T
	ret := db.Where("id = ?", id).Limit(1).Find(&s)
	if ret.Error != nil {
		return false, errs.Wrap("[database.ExistById]データ取得エラー", ret.Error)
	}
	return len(s) > 0, nil
}

// GetById は指定IDのレコードを返却します
func GetById[T any](ctx context.Context, db *gorm.DB, id uint32) (*T, error) {
	var s T
	ret := db.Where("id = ?", id).Limit(1).Find(&s)
	if ret.Error != nil {
		return nil, errs.Wrap("[database.GetById]データ取得エラー", ret.Error)
	}
	if ret.RowsAffected == 0 {
		return nil, errRecordNotFound
	}
	return &s, nil
}
