package content

import (
	"gitlab.kaonavi.jp/ae/sardine/internal/apps/setting/content/model"
)

type UrlService interface {
	// NewValidatedUrl はinputを検証して、検証済みのmodelを返却します
	NewValidatedUrl(UrlInput) (*model.ValidatedUrl, error)
}
