package database

import (
	"context"
	"errors"

	"gitlab.kaonavi.jp/ae/sardine/internal/errs"
	"gorm.io/gorm"
)

var (
	ErrRecordNotFound = errs.NewNotFound("record not found")
)

// Get は単一レコードの検索を行った結果を返します
func Get[T any](_ context.Context, db *gorm.DB) (*T, error) {
	var s T
	if err := db.First(&s).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrRecordNotFound
		}
		return nil, errs.NewInternalError("failed to database.Get: %v", err)
	}
	return &s, nil
}

// GetAll は複数レコードの検索を行った結果を返します
func GetAll[T any](_ context.Context, db *gorm.DB) ([]T, error) {
	var s []T
	if err := db.Find(&s).Error; err != nil {
		return nil, errs.NewInternalError("failed to database.GetAll: %v", err)
	}
	return s, nil
}

// Exist は検索を行った結果が存在する場合はnil、存在しない場合はerrorを返却します
func Exist[T any](ctx context.Context, db *gorm.DB) error {
	_, err := Get[T](ctx, db)
	return err
}

// ExistById は指定IDのレコードが存在する場合はnil、存在しない場合はerrorを返却します
func ExistById[T any](ctx context.Context, db *gorm.DB, id uint32) error {
	_, err := GetById[T](ctx, db, id)
	return err
}

// GetById は指定IDのレコードを返却します
func GetById[T any](ctx context.Context, db *gorm.DB, id uint32) (*T, error) {
	query := db.Where("id = ?", id)
	return Get[T](ctx, query)
}
