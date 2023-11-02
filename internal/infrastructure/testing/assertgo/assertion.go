package assertgo

type Assertion interface {
	WithTypeKeyManipulation(manipulationFunc func(string) string) Assertion
	IsEqualTo(target interface{})
}
