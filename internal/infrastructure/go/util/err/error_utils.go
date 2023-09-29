package err

func PanicIfError(err error) {
	PanicIfError1("placeholder", err)
}

func PanicIfError1[T any](result T, err error) T {
	if err != nil {
		panic(err)
	}
	return result
}
