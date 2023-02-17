package service

import (
	"net/url"

	"gitlab.kaonavi.jp/ae/sardine/internal/apps/setting/content"
	"gitlab.kaonavi.jp/ae/sardine/internal/apps/setting/content/model"
	"gitlab.kaonavi.jp/ae/sardine/internal/errs"
	"gitlab.kaonavi.jp/ae/sardine/internal/utils/validate"
)

var (
	allowDomains = map[string]struct{}{
		"www.youtube.com": {},
		"m.youtube.com":   {},
		"youtu.be":        {},
	}
)

func NewUrl() content.UrlService {
	return &urlService{}
}

type urlService struct{}

func (*urlService) NewValidatedUrl(in content.UrlInput) (*model.ValidatedUrl, error) {
	ers := errs.NewErrors()
	ers.AddError(validate.StringRequired("タイトル", &in.Title, 50))
	if err := validate.StringRequired("URL", &in.Url, 255); err != nil {
		ers.AddError(err)
		return nil, ers
	}

	parsed, err := url.Parse(in.Url)
	if err != nil || (parsed.Scheme != "http" && parsed.Scheme != "https") {
		ers.Add("URLの形式が不正です")
		return nil, ers
	}

	// 変なポート指定されるのもいやなので、ポート込みで見るためにHostを使っています
	// Hostname()を使うとポート抜きの文字列が帰ってくるので
	if _, ok := allowDomains[parsed.Host]; !ok {
		ers.Add("URLに許可されていないドメインが指定されています")
		return nil, ers
	}

	return errs.ErrorsOrNilWithValue(
		model.ValidatedUrl{Title: in.Title, Url: in.Url},
		ers,
	)
}
