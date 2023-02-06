package option

func Some[T any](v T) Option[T] {
	return Option[T]{Value: v, none: false}
}

func None[T any]() Option[T] {
	return Option[T]{none: true}
}

type Option[T any] struct {
	Value T
	none  bool
}

func (o Option[T]) None() bool {
	return o.none
}

func (o Option[T]) Some() (T, bool) {
	return o.Value, o.none
}

func (o Option[T]) SomeOr(def T) T {
	if o.none {
		return def
	}

	return o.Value
}
