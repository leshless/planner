package stupid

func NewReflect[T any](value T) func() T {
	return func() T {
		return value
	}
}
