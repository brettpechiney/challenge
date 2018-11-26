package challenge

import (
	"github.com/pkg/errors"

	"github.com/brettpechiney/challenge/cockroach"
)

// AttestationRepository provides an interface to extract information
// about an Attestation.
type AttestationRepository interface {
	Insert(attestation *NewAttestation) (string, error)
	Find(id string) (*Attestation, error)
	FindByUser(fname string, lname string) ([]*UserAttestation, error)
}

// attestationRepo abstracts the DAO and implements the
// AttestationRepository interface.
type attestationRepo struct {
	dao *cockroach.DAO
}

// NewAttestationRepo creates an attestationRepo.
func NewAttestationRepo(dao *cockroach.DAO) AttestationRepository {
	return attestationRepo{dao}
}

// Insert adds an Attestation to the database.
func (r attestationRepo) Insert(attestation *NewAttestation) (string, error) {
	const Query = `
		INSERT INTO attestation (
			claimant_id,
			attestor_id,
			claim
		) VALUES (
			$1,
			$2,
			$3
		)
		RETURNING id;`
	var id string
	err := r.dao.QueryRow(
		Query,
		attestation.ClaimantID,
		attestation.AttestorID,
		attestation.Claim,
	).Scan(&id)
	if err != nil {
		return "", errors.Wrapf(err, "failed to insert attestation")
	}
	return id, nil
}

// Find retrieves an Attestation by id.
func (r attestationRepo) Find(id string) (*Attestation, error) {
	const Query = `
		SELECT id,
			   claimant_id,
			   attestor_id,
			   claim
		FROM   attestation
		WHERE  id = $1;`
	var a Attestation
	err := r.dao.
		QueryRow(Query, id).
		Scan(
			&a.ID,
			&a.ClaimantID,
			&a.AttestorID,
			&a.Claim,
		)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to retrieve attestation")
	}
	return &a, nil
}

// FindByUser returns all Attestation claims and their attestors that belong to
// the user with the specified first and last name.
func (r attestationRepo) FindByUser(fname string, lname string) ([]*UserAttestation, error) {
	const Query = `
		WITH attestations AS (
		  SELECT a.id,
				 a.claimant_id,
				 a.attestor_id,
				 a.claim
		FROM   attestation AS a
			   JOIN challenge_user AS u
				 ON u.first_name = $1
					AND u.last_name = $2
		)
		SELECT a.claim,
			   concat_ws(' ', us.first_name, us.last_name) AS attestor_name
		FROM   attestations AS a
			   JOIN challenge_user AS us
				 ON us.id = a.attestor_id;`
	rows, err := r.dao.Query(Query, fname, lname)
	defer rows.Close()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to retrieve user attestations")
	}

	var attestations []*UserAttestation
	for rows.Next() {
		var a UserAttestation
		err := rows.Scan(&a.Claim, &a.AttestorName)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to scan user attestation")
		}
		attestations = append(attestations, &a)
	}
	return attestations, nil
}
