package models

// Attestation represents a verified claim.
type Attestation struct {
	ID         string `json:"id" sql:",type:uuid"`
	ClaimantID string `json:"claimantId" sql:",type:uuid"`
	AttestorID string `json:"attestorId" sql:",type:uuid"`
	Claim      string `json:"claim"`
}

// NewAttestation is an Attestation that has not yet been saved
// to the database.
type NewAttestation struct {
	ClaimantID string `json:"claimantId" sql:",type:uuid"`
	AttestorID string `json:"attestorId" sql:",type:uuid"`
	Claim      string `json:"claim"`
}

// OK validates the fields on a NewUser.
func (u *NewAttestation) OK() error {
	if len(u.ClaimantID) == 0 {
		return errMissingField("ClaimantID")
	}
	if len(u.AttestorID) == 0 {
		return errMissingField("AttestorID")
	}
	if len(u.Claim) == 0 {
		return errMissingField("Claim")
	}
	if len(u.Claim) > 100 {
		return &errMaxLengthExceeded{"Claim", 100}
	}
	return nil
}
