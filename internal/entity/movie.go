package entity

import (
	"time"
)

const (
	// MovieConvertStatusBeforeUpload 動画変換ステータス - ファイルアップロード前
	MovieConvertStatusBeforeUpload uint8 = 1
	// MovieConvertStatusUploaded 動画変換ステータス - 動画ファイルアップロード完了
	MovieConvertStatusUploaded uint8 = 2
	// MovieConvertStatusConverting 動画変換ステータス - 動画変換中
	MovieConvertStatusConverting uint8 = 3
	// MovieConvertStatusConverted 動画変換ステータス - 動画変換完了
	MovieConvertStatusConverted uint8 = 4
	// MovieConvertStatusConvertError 動画変換ステータス - 動画変換エラー
	MovieConvertStatusConvertError uint8 = 9
)

type Movie struct {
	Id                        uint32
	ContentId                 uint32
	DeliveryFileName          string
	OriginalFileName          string
	ThumbnailDeliveryFileName *string
	Duration                  uint32
	ConvertStatus             uint8
	ConvertErrorDetail        *string
	CreatedAt                 time.Time
	CreatedBy                 uint32
	UpdatedAt                 time.Time
	UpdatedBy                 uint32
}

type Movies = []Movie
