package course

type UpdateELearningInput struct {
	Title                  string                `json:"title"`
	Description            *string               `json:"description"`
	Thumbnail              *UpdateThumbnailInput `json:"thumbnail"`
	IsRemoveThumbnailImage bool                  `json:"isRemoveThumbnailImage"`
	IsRequired             bool                  `json:"isRequired"`
	CategoryId             *uint32               `json:"categoryId"`
	From                   *string               `json:"from"`
	To                     *string               `json:"to"`
}

type UpdateThumbnailInput struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}
