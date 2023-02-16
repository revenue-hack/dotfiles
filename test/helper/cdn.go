package helper

import (
	"fmt"
	"os"
)

// GetThumbnailDeliveryUrl は配信用のURLを返却します
func GetThumbnailDeliveryUrl(courseId uint32, fileName string) string {
	return fmt.Sprintf(
		"%s/%s/%d/%s",
		os.Getenv("AWS_CDN_BASE_URL"),
		dbName,
		courseId,
		fileName,
	)
}
