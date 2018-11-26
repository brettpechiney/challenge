package challenge

// A UserAttestation shows an attestation claim for a user as well as the
// full name of the attestation authority backing it.
type UserAttestation struct {
	Claim        string `json:"claim"`
	AttestorName string `json:"attestorName"`
}
