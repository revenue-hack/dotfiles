package entity

import (
	"time"
)

const (
	// ContentTypeMovie コンテンツの種別 - 動画
	ContentTypeMovie uint8 = 1
	// ContentTypeFile コンテンツの種別 - ファイル
	ContentTypeFile uint8 = 2
	// ContentTypeUrl コンテンツの種別 - 外部URL
	ContentTypeUrl uint8 = 3
)

type Content struct {
	Id           uint32
	CourseId     uint32
	ContentType  uint8
	DisplayOrder uint16
	CreatedAt    time.Time
	CreatedBy    uint32
	UpdatedAt    time.Time
	UpdatedBy    uint32

	// relations

	Movie Movie
	File  File
	Url   Url
}

type Contents = []Content

// IsMovie は動画コンテンツの場合にtrueを返します
func (c *Content) IsMovie() bool {
	return c.ContentType == ContentTypeMovie
}

// IsFile はファイルコンテンツの場合にtrueを返します
func (c *Content) IsFile() bool {
	return c.ContentType == ContentTypeFile
}

// IsUrl は外部URLコンテンツの場合にtrueを返します
func (c *Content) IsUrl() bool {
	return c.ContentType == ContentTypeUrl
}