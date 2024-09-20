package mapster

func Filter[T any](array []T, fn func(T any) bool) []T {
	result := []T{}
	for _, item := range array {
		if fn(item) {
			result = append(result, item)
		}
	}
	return result
}

func Map[T any, TResult any](array []T, fn func(t T) TResult) []TResult {
	result := []TResult{}
	for _, item := range array {
		result = append(result, fn(item))
	}
	return result
}
