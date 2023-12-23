package lambda

type Validatable interface {
	Validate() error
}
