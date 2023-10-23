package pkg

import "fmt"

func Pagination(limit int, page int, totalData int) string {
	defaultLimit := 100
	defaultoffset := 0

	if limit != 0 && limit <= defaultLimit {
		defaultLimit = limit
	}

	if page != 0 {
		defaultoffset = (page * defaultLimit) - defaultLimit
	}

	query := fmt.Sprintf("LIMIT %d OFFSET %d", defaultLimit, defaultoffset)

	return query
}