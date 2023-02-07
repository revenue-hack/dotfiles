package search

type Input struct {
	CategoryId    *uint32 `json:"categoryId"`
	SearchWord    *string `json:"searchWord"`
	NextPageToken *string `json:"nextPageToken"`
}
