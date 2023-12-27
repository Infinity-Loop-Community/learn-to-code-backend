package err

import "fmt"

func PanicIfNil[T any](name string, target *T) *T {
	if target == nil {
		panic(fmt.Errorf("%s is nil", name))
	}

	return target
}

func PanicIfError(err error) {
	PanicIfError1("placeholder", err)
}

func PanicIfError1[T any](result T, err error) T {
	if err != nil {
		panic(err)
	}
	return result
}
