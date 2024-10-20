package kickcom

func ptr[T any](in T) *T {
	return &in
}
