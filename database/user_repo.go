package database

import (
	"github.com/brettpechiney/challenge/models"
)

// UserRepository provides an interface to extract information about a User.
type UserRepository interface {
	Insert(user *models.NewUser) (string, error)
	Find(id string) (*models.User, error)
	FindAll() ([]*models.User, error)
	Update(user *models.User) (*models.User, error)
}

// UserRepo abstracts the DAO and implements the UserRepository interface.
type UserRepo struct {
	dao *DAO
}

// NewUserRepo creates a UserRepo.
func NewUserRepo(dao *DAO) *UserRepo {
	return &UserRepo{dao}
}

// Insert adds a User to the database.
func (r *UserRepo) Insert(user *models.NewUser) (string, error) {
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
		return "", err
	}
	return id, nil
}

// Find retrieves a user by id.
func (r *UserRepo) Find(id string) (*models.User, error) {
	const Query = `
		SELECT id,
			   first_name,
			   last_name,
			   username,
			   role,
			   last_login
		FROM   challenge_user
		WHERE  id = $1;`
	var u models.User
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
		return nil, err
	}
	return &u, nil
}

// FindAll retrieves all users.
func (r *UserRepo) FindAll() ([]*models.User, error) {
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
		return nil, err
	}

	var users []*models.User
	for rows.Next() {
		var u models.User
		err := rows.Scan(
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
		users = append(users, &u)
	}
	return users, nil
}

// Update makes changes to a user.
func (r *UserRepo) Update(user *models.User) (*models.User, error) {
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
	var u models.User
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
