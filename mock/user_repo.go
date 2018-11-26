package mock

import "github.com/brettpechiney/challenge/challenge"

// UserRepo is a mock implementation of challenge.UserRepo
type UserRepo struct {
	InsertFn       func(u *challenge.NewUser) (string, error)
	InsertInvoked  bool
	FindFn         func(id string) (*challenge.User, error)
	FindInvoked    bool
	FindAllFn      func() ([]*challenge.User, error)
	FindAllInvoked bool
	UpdateFn       func(u *challenge.User) (*challenge.User, error)
	UpdateInvoked  bool
}

func (r *UserRepo) Insert(u *challenge.NewUser) (string, error) {
	r.InsertInvoked = true
	return r.InsertFn(u)
}

func (r *UserRepo) Find(id string) (*challenge.User, error) {
	r.FindInvoked = true
	return r.FindFn(id)
}

func (r *UserRepo) FindAll() ([]*challenge.User, error) {
	r.FindAllInvoked = true
	return r.FindAllFn()
}

func (r *UserRepo) Update(u *challenge.User) (*challenge.User, error) {
	r.UpdateInvoked = true
	return r.UpdateFn(u)
}
