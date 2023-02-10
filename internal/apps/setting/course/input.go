package course

type UpdateELearningInput struct {
	Title          string  `json:"title"`
	Description    *string `json:"description"`
	ThumbnailImage *string `json:"thumbnailImage"`
	IsRequired     bool    `json:"isRequired"`
	CategoryId     *uint32 `json:"categoryId"`
	From           *string `json:"from"`
	To             *string `json:"to"`
}
