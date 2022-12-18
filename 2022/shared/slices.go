package shared

func MakeSliceInit[T comparable](len int, defaultValue T) (res []T) {
	res = make([]T, len)
	for i := 0; i < len; i++ {
		res[i] = defaultValue
	}
	return res
}

func SliceConvert[TIn comparable, TOut comparable](slice []TIn, converter func(in TIn) (TOut, error)) (res []TOut) {
	res = make([]TOut, len(slice))
	for i, v := range slice {
		if cv, err := converter(v); err == nil {
			res[i] = cv
		}
	}
	return res
}
