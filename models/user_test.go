package models_test

import (
	"fmt"
	"testing"

	"github.com/brettpechiney/challenge/models"
)

func TestNewUserValidation(t *testing.T) {
	const FirstName = "FirstName"
	const LastName = "LastName"
	const Username = "Username"
	const Role = "customer"
	const LongString string = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	testCases := []struct {
		Name      string
		TestModel *models.NewUser
		ShouldErr bool
	}{
		{
			"ValidModel",
			&models.NewUser{
				FirstName: FirstName,
				LastName:  LastName,
				Username:  Username,
				Role:      Role,
			},
			false,
		},
		{
			"MissingFirstName",
			&models.NewUser{
				FirstName: "",
				LastName:  LastName,
				Username:  Username,
				Role:      Role,
			},
			true,
		},
		{
			"FirstNameLengthGreaterThan50",
			&models.NewUser{
				FirstName: LongString,
				LastName:  LastName,
				Username:  Username,
				Role:      Role,
			},
			true,
		},
		{
			"MissingLastName",
			&models.NewUser{
				FirstName: FirstName,
				Username:  Username,
				Role:      Role,
			},
			true,
		},
		{
			"LastNameLengthGreaterThan50",
			&models.NewUser{
				FirstName: FirstName,
				LastName:  LongString,
				Username:  Username,
				Role:      Role,
			},
			true,
		},
		{
			"MissingUSername",
			&models.NewUser{
				FirstName: FirstName,
				LastName:  LastName,
				Role:      Role,
			},
			true,
		},
		{
			"UsernameLengthGreaterThan50",
			&models.NewUser{
				FirstName: FirstName,
				LastName:  LastName,
				Username:  LongString,
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
