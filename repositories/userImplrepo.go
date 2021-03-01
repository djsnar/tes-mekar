package repositories

import (
	"database/sql"
	"fmt"
	"mekar/model"
	"mekar/utils"
)

// type UserRepoImpl
type UserRepoImpl struct {
	db *sql.DB
}

// func GetAllUser
func (s *UserRepoImpl) GetAllUser() ([]*model.User, error) {
	var dataUser []*model.User
	query := fmt.Sprintf(utils.SELECT_ALL_USER)
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		user := model.User{}
		var err = rows.Scan(&user.NamaUser, &user.TanggalLahir, &user.NoKTP, &user.Pekerjaan, &user.Pendidikan)
		if err != nil {
			return nil, err
		}
		dataUser = append(dataUser, &user)
	}
	return dataUser, nil
}

// func GetUserByID
func (s *UserRepoImpl) GetUserByID(NoKTP string) (model.User, error) {
	var user model.User
	query := utils.SELECT_USER_BY_ID
	err := s.db.QueryRow(query, NoKTP).Scan(&user.NamaUser, &user.TanggalLahir, &user.NoKTP, &user.Pekerjaan, &user.Pendidikan)
	return user, err
}

// func CreateNewUser
func (s *UserRepoImpl) CreateNewUser(user model.User) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(utils.INSERT_USER)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(user.NamaUser, user.TanggalLahir, user.NoKTP, user.Pekerjaan, user.Pendidikan)
	if err != nil {
		return tx.Rollback()
	}

	stmt.Close()
	return tx.Commit()
}

// func UpdateUser
func (s *UserRepoImpl) UpdateUser(NoKTP string, user *model.User) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(utils.UPDATE_USER)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user.NamaUser, user.TanggalLahir, user.NoKTP, user.Pekerjaan, user.Pendidikan)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	stmt.Close()
	return err
}

//func DeleteUser
func (s *UserRepoImpl) DeleteUser(NoKTP string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(utils.DELETE_USER)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(NoKTP)
	if err != nil {
		return tx.Rollback()
	}

	stmt.Close()
	return tx.Commit()
}

//func InitUserRepoImpl
func InitUserRepoImpl(db *sql.DB) UserRepository {
	return &UserRepoImpl{db}
}
