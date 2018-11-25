package models_test

import (
	"fmt"
	"github.com/brettpechiney/challenge/models"
	"math"
	"strings"
	"testing"
)

func TestNewAttestationValidation(t *testing.T) {
	const ClaimantID = "testID"
	const AttestorID = "testID"
	const Claim = "testClaim"
	longString := getLongString(100)
	testCases := []struct {
		Name      string
		TestModel *models.NewAttestation
		ShouldErr bool
	}{
		{
			"ValidModel",
			&models.NewAttestation{
				ClaimantID: ClaimantID,
				AttestorID: AttestorID,
				Claim:      Claim,
			},
			false,
		},
		{
			"MissingClaimantID",
			&models.NewAttestation{
				AttestorID: AttestorID,
				Claim:      Claim,
			},
			true,
		},
		{
			"MissingAttestorID",
			&models.NewAttestation{
				ClaimantID: ClaimantID,
				Claim:      Claim,
			},
			true,
		},
		{
			"MissingClaim",
			&models.NewAttestation{
				ClaimantID: ClaimantID,
			},
			true,
		},
		{
			"ClaimLengthGreaterThan100",
			&models.NewAttestation{
				ClaimantID: ClaimantID,
				Claim:      longString,
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

func getLongString(len int) string {
	const str = "test"
	floatIterations := float64(len / 4) // Divide by length of str
	iterations := int(math.Ceil(floatIterations))

	var sb strings.Builder
	for i := 0; i <= iterations; i++ {
		sb.WriteString(str)
	}
	return sb.String()
}
