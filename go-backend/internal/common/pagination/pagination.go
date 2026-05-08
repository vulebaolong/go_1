package pagination

import "strconv"

type Query struct {
	Page     int
	PageSize int
	Offset   int
}

func Get(pageString string, pageSizeString string) Query {
	page, err := strconv.Atoi(pageString)
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeString)
	if err != nil || pageSize <= 0 {
		pageSize = 3
	}

	offset := (page - 1) * pageSize

	return Query{
		Page:     page,
		PageSize: pageSize,
		Offset:   offset,
	}
}
