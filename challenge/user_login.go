package challenge

// A UserLogin represents a request to log in to the application.
type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// OK validates the fields on a UserLogin
func (ul UserLogin) OK() error {
	if len(ul.Username) == 0 {
		return errMissingField("Username")
	}
	if len(ul.Username) > 50 {
		return &errMaxLengthExceeded{"Username", 50}
	}
	if len(ul.Password) == 0 {
		return errMissingField("Password")
	}
	return nil
}
