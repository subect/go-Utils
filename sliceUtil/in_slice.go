package sliceUtil

// InSlice 判断value是否在slice中
func InSlice[T comparable](value T, slice []T) bool {
	for _, v := range slice {
		if value == v {
			return true
		}
	}
	return false
}
