package challenge

// Okay is an interface implemented by types that can validate
// themselves.
type Okay interface {
	OK() error
}
