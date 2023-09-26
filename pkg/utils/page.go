package utils

func CalculateLimitOffset(page int, perPage int) (int, int) {
	return perPage, (page - 1) * perPage
}
