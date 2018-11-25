package database

import (
	"github.com/brettpechiney/challenge/models"
)

// AttestationRepository provides an interface to extract information
// about an Attestation.
type AttestationRepository interface {
	Insert(attestation models.NewAttestation) (string, error)
	Find(id string) (*models.Attestation, error)
	FindByUser(fname string, lname string) ([]*models.UserAttestation, error)
}

// AttestationRepo abstracts the DAO and implements the
// AttestationRepository interface.
type AttestationRepo struct {
	dao *DAO
}

// NewAttestationRepo creates an AttestationRepo.
func NewAttestationRepo(dao *DAO) *AttestationRepo {
	return &AttestationRepo{dao}
}

// Insert adds an Attestation to the database.
func (r *AttestationRepo) Insert(attestation *models.NewAttestation) (string, error) {
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
		return "", err
	}
	return id, nil
}

// Find retrieves an Attestation by id.
func (r *AttestationRepo) Find(id string) (*models.Attestation, error) {
	const Query = `
		SELECT id,
			   claimant_id,
			   attestor_id,
			   claim
		FROM   attestation
		WHERE  id = $1;`
	var a models.Attestation
	err := r.dao.
		QueryRow(Query, id).
		Scan(
			&a.ID,
			&a.ClaimantID,
			&a.AttestorID,
			&a.Claim,
		)
	if err != nil {
		return nil, err
	}
	return &a, nil
}

// FindByUser returns all Attestation claims and their attestors that belong to
// the user with the specified first and last name.
func (r *AttestationRepo) FindByUser(fname string, lname string) ([]*models.UserAttestation, error) {
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
		return nil, err
	}

	var attestations []*models.UserAttestation
	for rows.Next() {
		var a models.UserAttestation
		err := rows.Scan(&a.Claim, &a.AttestorName)
		if err != nil {
			return nil, err
		}
		attestations = append(attestations, &a)
	}
	return attestations, nil
}
