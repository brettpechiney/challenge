package challenge

import "time"

// A User represents an end-user of Challenge.
type User struct {
	ID        string    `json:"id" sql:",type:uuid"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	LastLogin time.Time `json:"lastLogin"`
}

// A NewUser represents an end-user of Challenge that has not yet been
// saved to the database.
type NewUser struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Role      string `json:"role"`
}

var validRoles = map[string]bool{
	"customer":              true,
	"administrator":         true,
	"attestation authority": true,
}

// OK validates the fields on a NewUser.
func (u NewUser) OK() error {
	if len(u.FirstName) == 0 {
		return errMissingField("FirstName")
	}
	if len(u.FirstName) > 50 {
		return &errMaxLengthExceeded{"FirstName", 50}
	}
	if len(u.LastName) == 0 {
		return errMissingField("LastName")
	}
	if len(u.LastName) > 50 {
		return &errMaxLengthExceeded{"LastName", 50}
	}
	if len(u.Username) == 0 {
		return errMissingField("Username")
	}
	if len(u.Username) > 50 {
		return &errMaxLengthExceeded{"Username", 50}
	}
	if len(u.Password) == 0 {
		return errMissingField("Password")
	}
	if len(u.Role) == 0 {
		return errMissingField("Role")
	}
	if !validRoles[u.Role] {
		return &errInvalidValue{
			"Role",
			[]string{
				"customer",
				"administrator",
				"attestation authority",
			},
		}
	}
	return nil
}
