package paging

type Paginator struct {
	PageIndex int
	PageSize  int
	Pages     int
	Total     int
	Items     interface{}
}

func NewPaginator(pageIndex, pageSize, total int, items interface{}) *Paginator {
	pages := getPages(total, pageSize)

	return &Paginator{
		PageIndex: pageIndex,
		PageSize:  pageSize,
		Pages:     pages,
		Total:     total,
		Items:     items,
	}
}

func getPages(total, pageSize int) int {
	return (total + pageSize - 1) / pageSize
}

// GetOffset Offset sets the first node to return from the query
func GetOffset(pageIndex, pageSize int) int {
	return (pageIndex - 1) * pageSize
}
