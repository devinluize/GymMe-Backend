package responses

type Pagination struct {
	CurrentPage int         `json:"current_page"`
	TotalPages  int         `json:"total_pages"`
	PageSize    int         `json:"page_size"`
	TotalItems  int         `json:"total_items"`
	Items       interface{} `json:"items"`
}

// Example of a fu nction that creates a Pagination struct
func NewPagination(currentPage, pageSize, totalItems int, items interface{}) *Pagination {
	totalPages := (totalItems + pageSize - 1) / pageSize

	return &Pagination{
		CurrentPage: currentPage,
		TotalPages:  totalPages,
		PageSize:    pageSize,
		TotalItems:  totalItems,
		Items:       items,
	}
}
