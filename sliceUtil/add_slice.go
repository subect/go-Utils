package sliceUtil

// AddIndex 在index处添加元素
func AddIndex[T comparable](src []T, element T, index int) ([]T, error) {
	length := len(src)
	if index < 0 || index >= length {
		return nil, NewErrIndexOutOfRange(length, index)
	}

	//先将src扩展一个元素
	var zeroValue T
	src = append(src, zeroValue)
	for i := len(src) - 1; i > index; i-- {
		if i-1 >= 0 {
			src[i] = src[i-1]
		}
	}
	src[index] = element
	return src, nil
}
