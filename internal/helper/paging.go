package helper

import (
	"gokomodo_test/internal/consts"
	"math"
)

// LimitDefaultValue set default value limit
func LimitDefaultValue(origin int) int {
	if origin < 1 {
		return consts.PagingLimitDefaultValue
	}

	if origin > consts.PagingMaxLimit {
		return consts.PagingMaxLimit
	}

	return origin
}

// PageDefaultValue set default value page
func PageDefaultValue(origin int) int {
	if origin < 1 {
		return consts.PagingPageDefaultValue
	}

	return origin
}

// PageCalculate calculate total page from count
func PageCalculate(count int, limit int) int {
	if count <= limit {
		return 1
	}

	return int(math.Ceil(float64(count) / float64((limit))))
}

// OffsetDefaultValue set default offset
func OffsetDefaultValue(page int, limit int) int {
	if page < 1 {
		return consts.PagingOffsetDefaultValue
	}

	return (page - 1) * limit
}

// PageToOffset calculate
func PageToOffset(limit, page int) int {
	if page <= 0 {
		return 0
	}

	return (page - 1) * limit
}
