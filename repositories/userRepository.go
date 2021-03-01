package repositories

import "mekar/model"

// type UserRepository
type UserRepository interface {
	GetAllUser() ([]*model.User, error)
	GetUserByID(NoKTP string) (model.User, error)
	CreateNewUser(user model.User) error
	UpdateUser(NoKTP string, user *model.User) error
	DeleteUser(NoKTP string) error
}
