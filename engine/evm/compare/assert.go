package compare

type Constructor[expectedT any] func(raw string) (Compare[expectedT], error)

type Compare[expectedT any] interface {
	Fixed(actual expectedT) bool
	Min(actual expectedT) bool
	Max(actual expectedT) bool
	Magic(actual expectedT) bool
}

// Falsy embed to rewrite only required compare methods: for example, min/max for address must be false
type Falsy[T any] struct{}

func (f *Falsy[T]) Fixed(_ T) bool {
	return false
}
func (f *Falsy[T]) Min(_ T) bool {
	return false
}
func (f *Falsy[T]) Max(_ T) bool {
	return false
}
func (f *Falsy[T]) Magic(_ T) bool {
	return false
}
