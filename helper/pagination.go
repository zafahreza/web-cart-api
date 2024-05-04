package helper

import (
	"strconv"
	"web-cart-api/model"
)

func SetPaginationFromQuery(queryLimit string, queryPage string) model.QueryPagination {

	limit := -1
	page := -1
	offset := -1

	if queryLimit != "" {
		limit, _ = strconv.Atoi(queryLimit)
	}

	if queryPage != "" {
		page, _ = strconv.Atoi(queryPage)
	}

	if limit != -1 || page != -1 {
		offset = (page - 1) * limit
	}

	pagination := model.QueryPagination{
		Limit:  limit,
		Offset: offset,
	}

	return pagination
}
