package svgpath

func reduce[T, M any](s []T, f func(M, T) M, initValue M) M {
	acc := initValue
	for _, v := range s {
		acc = f(acc, v)
	}
	return acc
}

func reduceE[T, M any](s []T, f func(M, T) (M, error), initValue M) (acc M, err error) {
	acc = initValue
	for _, v := range s {
		acc, err = f(acc, v)
		if err != nil {
			return
		}
	}

	return
}
