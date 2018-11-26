package mock

import "github.com/brettpechiney/challenge/challenge"

// AttestationRepo is a mock implementation of challenge.AttestationRepo
type AttestationRepo struct {
	InsertFn          func(a *challenge.NewAttestation) (string, error)
	InsertInvoked     bool
	FindFn            func(id string) (*challenge.Attestation, error)
	FindInvoked       bool
	FindByUserFn      func(fname string, lname string) ([]*challenge.UserAttestation, error)
	FindByUserInvoked bool
}

func (r *AttestationRepo) Insert(a *challenge.NewAttestation) (string, error) {
	r.InsertInvoked = true
	return r.InsertFn(a)
}

func (r *AttestationRepo) Find(id string) (*challenge.Attestation, error) {
	r.FindInvoked = true
	return r.FindFn(id)
}

func (r *AttestationRepo) FindByUser(fname string, lname string) ([]*challenge.UserAttestation, error) {
	r.FindByUserInvoked = true
	return r.FindByUserFn(fname, lname)
}
