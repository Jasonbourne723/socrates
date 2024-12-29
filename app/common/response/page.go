package response

type Page[T interface{}] struct {
	PageIndex  int32
	PageSize   int32
	TotalCount int64
	TotalPage  int64
	Rows       []T
}
