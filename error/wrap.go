package error

func WrapIntErr(i int, e error) int {
	if e != nil {
		panic(e)
	}

	return i
}
