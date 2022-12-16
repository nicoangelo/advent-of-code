package shared

func MakeSliceInit[T comparable](len int, defaultValue T) (res []T) {
	res = make([]T, len)
	for i := 0; i < len; i++ {
		res[i] = defaultValue
	}
	return res
}
