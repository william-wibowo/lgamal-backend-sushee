package repository

import (
	"final-project-backend/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	GetUserById(i int) (*entity.User, error)
	GetUserByEmailOrUsername(string) (*entity.User, error)
	CheckEmailExistence(string) (int, error)
	CheckUsernameExistence(string) (int, error)
	AddNewUser(*entity.User) (*entity.User, error)
	AddNewUserSession(s *entity.Session) (*entity.Session, error)
	GetUserSessionByRefreshToken(t string) (*entity.Session, error) 
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

type UserRepositoryConfig struct {
	DB *gorm.DB
}

func NewUserRepository(c UserRepositoryConfig) UserRepository {
	return &UserRepositoryImpl{
		db: c.DB,
	}
}

func (r *UserRepositoryImpl) GetUserById(i int) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("id = ?", i).First(&user).Error
	return &user, err
}

func (r *UserRepositoryImpl) GetUserByEmailOrUsername(i string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("email = ?", i).Or("username = ?", i).First(&user).Error
	return &user, err
}

func (r *UserRepositoryImpl) CheckEmailExistence(i string) (int, error) {
	var user entity.User
	tx := r.db.Clauses(clause.OnConflict{DoNothing: true}).Where("email = ?", i).First(&user)
	return int(tx.RowsAffected), tx.Error
}

func (r *UserRepositoryImpl) CheckUsernameExistence(i string) (int, error) {
	var user entity.User
	tx := r.db.Clauses(clause.OnConflict{DoNothing: true}).Where("email = ?", i).Or("username = ?", i).First(&user)
	return int(tx.RowsAffected), tx.Error
}

func (r *UserRepositoryImpl) AddNewUser(u *entity.User) (*entity.User, error) {
	err := r.db.Create(&u).Error
	return u, err
}

func (r *UserRepositoryImpl) AddNewUserSession(s *entity.Session) (*entity.Session, error) {
	err := r.db.Create(&s).Error
	return s, err
}

func (r *UserRepositoryImpl) GetUserSessionByRefreshToken(t string) (*entity.Session, error) {
	var session entity.Session
	err := r.db.Where("refresh_token = ?", t).First(&session).Error
	return &session, err
} 

// func (r *UserRepositoryImpl) GetUserByEmail(email string) (*entity.User, error) {
// 	var user entity.User
// 	err := r.db.Where("email = ?", email).First(&user).Error
// 	return &user, err
// }

// func (r *UserRepositoryImpl) GetUserDetailById(id int) (*entity.User, error) {
// 	var user entity.User
// 	err := r.db.Model(&user).Preload("Wallet").Where("id = ?", id).First(&user).Error
// 	return &user, err
// }

// func (r *UserRepositoryImpl) UpdateUserPassword(userId int, newPass string) error {
// 	var user entity.User
// 	err := r.db.Model(&user).
// 		Where("id = ?", userId).
// 		Update("password", newPass).
// 		Error
// 	return err
// }
