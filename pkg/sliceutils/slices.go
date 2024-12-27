package sliceutils

import (
	"sort"
	"strings"
	"unicode"
)

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

func SliceAppendRange(s []int, start int, count int) []int {
	if s == nil {
		panic("Cannot append to nil slice")
	}
	for i := 0; i < count; i++ {
		s = append(s, start+i)
	}
	return s
}

func RemoveAtIndex(slice []int, index int) []int {
	slice_new := make([]int, len(slice))
	copy(slice_new, slice)
	return append(slice_new[:index], slice_new[index+1:]...)
}

func SortAlphabeticallyIgnoreCase(slice []string) {
	sort.Slice(slice, func(i, j int) bool {
		x := strings.Map(unicode.ToUpper, slice[i])
		y := strings.Map(unicode.ToUpper, slice[j])
		return x < y
	})
}
