package utils

func worker[T any, U any](c chan T, ret chan U, f func(T) U) {
	for el := range c {
		ret <- f(el)
	}
}

func Workers[T any, U any](params []T, f func(T) U, c int) []U {
	cParam := make(chan T)
	cResults := make(chan U)

	for i := 0; i < c; i++ {
		go worker(cParam, cResults, f)
	}

	for _, param := range params {
		cParam <- param
	}

	close(cParam)

	var ret []U
	for el := range cResults {
		ret = append(ret, el)
	}
	close(cResults)
	return ret

}
