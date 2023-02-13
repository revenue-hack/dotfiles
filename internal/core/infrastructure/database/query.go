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
	return err == errRecordNotFound
}

// Get は単一レコードの検索を行った結果を返します
func Get[T any](_ context.Context, db *gorm.DB) (*T, error) {
	var s T
	if err := db.First(&s).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errRecordNotFound
		}
		return nil, errs.Wrap("[database.Get]データ取得エラー", err)
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
func Exist[T any](ctx context.Context, db *gorm.DB) (bool, error) {
	var s T
	if err := db.First(&s).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, errs.Wrap("[database.Exist]データ取得エラー", err)
	}
	return true, nil
}

// ExistById は指定IDのレコードが存在する場合はnil、存在しない場合はerrorを返却します
func ExistById[T any](ctx context.Context, db *gorm.DB, id uint32) (bool, error) {
	query := db.Where("id = ?", id)
	var s T
	if err := query.First(&s).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, errs.Wrap("[database.ExistById]データ取得エラー", err)
	}
	return true, nil
}

// GetById は指定IDのレコードを返却します
func GetById[T any](ctx context.Context, db *gorm.DB, id uint32) (*T, error) {
	query := db.Where("id = ?", id)
	var s T
	if err := query.First(&s).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errRecordNotFound
		}
		return nil, errs.Wrap("[database.GetById]データ取得エラー", err)
	}
	return &s, nil
}
