package helper

import (
	"time"
)

var (
	// FixedTime はFixedMockTimeより過去日の固定の日付
	// 初期データとしてFixedTimeを設定し、処理後にFixedMockTimeの値に更新されていることを確認する場合に使用します
	FixedTime time.Time
	// FixedMockTime はtimer packageのモックが返却する固定の日付です
	FixedMockTime time.Time
	// Location はタイムゾーンです
	Location *time.Location
)

func init() {
	Location, _ = time.LoadLocation("Asia/Tokyo")
	FixedTime = time.Date(2023, 2, 9, 10, 0, 0, 0, Location)
	FixedMockTime = time.Date(2023, 4, 3, 12, 34, 56, 0, Location)
}
