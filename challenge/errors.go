package challenge

import (
	"fmt"
	"strings"
)

// errMissingField represents an error message indicating the absence
// of a required field.
type errMissingField string

// Error generates an errMissingField message.
func (e errMissingField) Error() string {
	return string(e) + " is required"
}

// errMaxLengthExceeded represents an error message indicating that
// a string field is too long.
type errMaxLengthExceeded struct {
	field  string
	maxLen int
}

// Error generates an errMaxLengthExceeded message.
func (e errMaxLengthExceeded) Error() string {
	return fmt.Sprintf("field %s exceeds maximum length of %d", e.field, e.maxLen)
}

// errInvalidValue represents an error message indicating that a
// value is not supported.
type errInvalidValue struct {
	field       string
	validValues []string
}

// Error generates an errInvalidValue message.
func (e errInvalidValue) Error() string {
	return fmt.Sprintf("field %s contains value not in %s", e.field, strings.Join(e.validValues, ", "))
}
