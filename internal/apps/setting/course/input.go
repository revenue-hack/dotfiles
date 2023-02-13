package course

type UpdateELearningInput struct {
	Title                  string  `json:"title"`
	Description            *string `json:"description"`
	ThumbnailImage         *string `json:"thumbnailImage"`
	IsRemoveThumbnailImage bool    `json:"isRemoveThumbnailImage"`
	IsRequired             bool    `json:"isRequired"`
	CategoryId             *uint32 `json:"categoryId"`
	From                   *string `json:"from"`
	To                     *string `json:"to"`
}
