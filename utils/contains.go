package utils

func Contains[T comparable](l []T, el T) bool {
	for _, v := range l {
		if v == el {
			return true
		}
	}
	return false
}

func NotContains[T comparable](l []T, el T) bool {
	return !Contains(l, el)
}
