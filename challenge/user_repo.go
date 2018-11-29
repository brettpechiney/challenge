package challenge

import (
	"github.com/pkg/errors"

	"github.com/brettpechiney/challenge/cockroach"
)

// UserRepository provides an interface to extract information about a User.
type UserRepository interface {
	Insert(user *NewUser) (string, error)
	Find(id string) (*User, error)
	FindAll() ([]*User, error)
	GetPassword(username string) (string, error)
	Update(user *User) (*User, error)
}

// userRepo abstracts the DAO and implements the UserRepository interface.
type userRepo struct {
	dao *cockroach.DAO
}

// NewUserRepo creates a userRepo.
func NewUserRepo(dao *cockroach.DAO) UserRepository {
	return userRepo{dao}
}

// Insert adds a User to the database.
func (r userRepo) Insert(user *NewUser) (string, error) {
	const Query = `
		INSERT INTO challenge_user (
			first_name,
    		last_name,
    		username,
    		password,
    		role
		) VALUES(
			$1,
    		$2,
    		$3,
    		$4,
    		$5
		)
		RETURNING id;`
	var id string
	err := r.dao.QueryRow(
		Query,
		user.FirstName,
		user.LastName,
		user.Username,
		user.Password,
		user.Role,
	).Scan(&id)
	if err != nil {
		return "", errors.Wrapf(err, "failed to insert user")
	}
	return id, nil
}

// Find retrieves a user by id.
func (r userRepo) Find(id string) (*User, error) {
	const Query = `
		SELECT id,
			   first_name,
			   last_name,
			   username,
			   role,
			   last_login
		FROM   challenge_user
		WHERE  id = $1;`
	var u User
	err := r.dao.
		QueryRow(Query, id).
		Scan(
			&u.ID,
			&u.FirstName,
			&u.LastName,
			&u.Username,
			&u.Role,
			&u.LastLogin,
		)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to retrieve user")
	}
	return &u, nil
}

// FindAll retrieves all users.
func (r userRepo) FindAll() ([]*User, error) {
	const Query = `
		SELECT id,
			   first_name,
			   last_name,
			   username,
			   role,
			   last_login
		FROM   challenge_user;`
	rows, err := r.dao.Query(Query)
	defer rows.Close()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to retrieve users")
	}

	var users []*User
	for rows.Next() {
		var u User
		err := rows.Scan(
			&u.ID,
			&u.FirstName,
			&u.LastName,
			&u.Username,
			&u.Role,
			&u.LastLogin,
		)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to scan user")
		}
		users = append(users, &u)
	}
	return users, nil
}

// GetPassword returns the password for the user with the specified username.
func (r userRepo) GetPassword(username string) (string, error) {
	const Query = `
		SELECT password
		FROM   challenge_user
		WHERE  username = $1;`
	var pw string
	err := r.dao.QueryRow(Query, username).Scan(&pw)
	if err != nil {
		return "", errors.Wrapf(err, "failed to retrieve password")
	}
	return pw, nil
}

// Update makes changes to a user.
func (r userRepo) Update(user *User) (*User, error) {
	const Query = `
		UPDATE    challenge_user
		SET       first_name = $1,
			      last_name = $2,
			      username = $3,
				  role = $4
		WHERE  	  id = $5
		RETURNING id,
				  first_name,
				  last_name,
				  username,
				  role,
				  last_login;`
	var u User
	err := r.dao.QueryRow(
		Query,
		user.FirstName,
		user.LastName,
		user.Username,
		user.Role,
		user.ID,
	).
		Scan(
			&u.ID,
			&u.FirstName,
			&u.LastName,
			&u.Username,
			&u.Role,
			&u.LastLogin,
		)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
