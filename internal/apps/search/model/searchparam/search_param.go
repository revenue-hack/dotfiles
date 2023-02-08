package searchparam

type SearchParam interface {
	CategoryId() *uint32
	SearchWord() *string
	NextCourseId() *uint32
	Limit() uint32
}
