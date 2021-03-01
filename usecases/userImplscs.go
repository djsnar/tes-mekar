package usecases

import (
	"mekar/model"
	"mekar/repositories"
	"mekar/utils"
)

//type UserUseCaseImpl
type UserUseCaseImpl struct {
	userRepo repositories.UserRepository
}

// Anonymous function GetAllUser
func (s UserUseCaseImpl) GetAllUser() ([]*model.User, error) {
	user, err := s.userRepo.GetAllUser()
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Anonymous function GetUserByID
func (s UserUseCaseImpl) GetUserByID(NoKTP string) (model.User, error) {
	user, err := s.userRepo.GetUserByID(NoKTP)

	return user, err
}

//Anonymous function CreateNewService
func (s UserUseCaseImpl) CreateNewUser(user model.User) error {
	err := utils.ValidationHandler(user.NamaUser, user.TanggalLahir, user.NoKTP, user.Pekerjaan, user.Pendidikan)
	if err != nil {
		return err
	}
	err = s.userRepo.CreateNewUser(user)
	if err != nil {
		return err
	}
	return err
}

// Anonymous function UpdateUser
func (s UserUseCaseImpl) UpdateUser(NoKTP string, user *model.User) error {
	err := s.userRepo.UpdateUser(NoKTP, user)
	if err != nil {
		return err
	}
	return nil
}

// Anonymous function DeleteUser
func (s UserUseCaseImpl) DeleteUser(NoKTP string) error {
	err := s.userRepo.DeleteUser(NoKTP)
	if err != nil {
		return err
	}
	return nil
}

// func InitServiceUseCase
func InitUserUseCase(userRepo repositories.UserRepository) UserUsecase {
	return &UserUseCaseImpl{userRepo}
}
