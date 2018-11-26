package challenge_test

import (
	"fmt"
	"math"
	"strings"
	"testing"

	"github.com/brettpechiney/challenge/challenge"
)

func TestNewAttestationValidation(t *testing.T) {
	const ClaimantID = "testID"
	const AttestorID = "testID"
	const Claim = "testClaim"
	longString := getLongString(100)
	testCases := []struct {
		Name      string
		TestModel *challenge.NewAttestation
		ShouldErr bool
	}{
		{
			"ValidModel",
			&challenge.NewAttestation{
				ClaimantID: ClaimantID,
				AttestorID: AttestorID,
				Claim:      Claim,
			},
			false,
		},
		{
			"MissingClaimantID",
			&challenge.NewAttestation{
				AttestorID: AttestorID,
				Claim:      Claim,
			},
			true,
		},
		{
			"MissingAttestorID",
			&challenge.NewAttestation{
				ClaimantID: ClaimantID,
				Claim:      Claim,
			},
			true,
		},
		{
			"MissingClaim",
			&challenge.NewAttestation{
				ClaimantID: ClaimantID,
			},
			true,
		},
		{
			"ClaimLengthGreaterThan100",
			&challenge.NewAttestation{
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
