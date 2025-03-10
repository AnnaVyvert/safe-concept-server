package middleware

type contextKey struct {
	name string
}

func (k *contextKey) String() string {
	return "middleware context value " + k.name
}
