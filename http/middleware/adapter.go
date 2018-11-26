package middleware

import "net/http"

// An Adapter takes in and returns an http.Handler. It is very useful when
// writing middleware.
type Adapter func(http.Handler) http.Handler

// Adapt wraps an http.Handler with a series of Adapter types, returning
// the result of the first adapter. Note that the adapters are applied in
// reverse of the specified order.
func Adapt(h http.Handler, adapters ...Adapter) http.Handler {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}
