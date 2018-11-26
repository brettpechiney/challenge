package challenge_test

import (
	"fmt"
	"testing"

	"github.com/brettpechiney/challenge/challenge"
)

func TestNewUserValidation(t *testing.T) {
	const FirstName = "FirstName"
	const LastName = "LastName"
	const Username = "Username"
	const Password = "Password"
	const Role = "customer"
	const LongString string = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	testCases := []struct {
		Name      string
		TestModel *challenge.NewUser
		ShouldErr bool
	}{
		{
			"ValidModel",
			&challenge.NewUser{
				FirstName: FirstName,
				LastName:  LastName,
				Username:  Username,
				Password:  Password,
				Role:      Role,
			},
			false,
		},
		{
			"MissingFirstName",
			&challenge.NewUser{
				FirstName: "",
				LastName:  LastName,
				Username:  Username,
				Password:  Password,
				Role:      Role,
			},
			true,
		},
		{
			"FirstNameLengthGreaterThan50",
			&challenge.NewUser{
				FirstName: LongString,
				LastName:  LastName,
				Username:  Username,
				Password:  Password,
				Role:      Role,
			},
			true,
		},
		{
			"MissingLastName",
			&challenge.NewUser{
				FirstName: FirstName,
				Username:  Username,
				Password:  Password,
				Role:      Role,
			},
			true,
		},
		{
			"LastNameLengthGreaterThan50",
			&challenge.NewUser{
				FirstName: FirstName,
				LastName:  LongString,
				Username:  Username,
				Password:  Password,
				Role:      Role,
			},
			true,
		},
		{
			"MissingUSername",
			&challenge.NewUser{
				FirstName: FirstName,
				LastName:  LastName,
				Password:  Password,
				Role:      Role,
			},
			true,
		},
		{
			"UsernameLengthGreaterThan50",
			&challenge.NewUser{
				FirstName: FirstName,
				LastName:  LastName,
				Username:  LongString,
				Password:  Password,
				Role:      Role,
			},
			true,
		},
		{
			"MissingPassword",
			&challenge.NewUser{
				FirstName: FirstName,
				LastName:  LastName,
				Username:  Username,
				Role:      Role,
			},
			true,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s", tc.Name), func(t *testing.T) {
			err := tc.TestModel.OK()
			failure := ((err == nil) && tc.ShouldErr) || ((err != nil) && !tc.ShouldErr)
			if failure {
				t.Error(err)
			}
		})
	}
}
