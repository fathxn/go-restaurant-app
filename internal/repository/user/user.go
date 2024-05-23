package user

import (
	"crypto/cipher"
	"go-restaurant-app/internal/model"
	"gorm.io/gorm"
)

type userRepo struct {
	db      *gorm.DB
	gcm     cipher.AEAD
	time    uint32
	memory  uint32
	threads uint8
	keyLen  uint32
}

func GetRepository(db *gorm.DB, time uint32, memory uint32, threads uint8, keyLen uint32) Repository {
	return &userRepo{
		db:      db,
		time:    time,
		memory:  memory,
		threads: threads,
		keyLen:  keyLen,
	}
}

func (ur userRepo) RegisterUser(userData model.User) (model.User, error) {
	if err := ur.db.Create(&userData).Error; err != nil {
		return model.User{}, err
	}
	return userData, nil
}

func (ur userRepo) CheckRegistered(username string) (bool, error) {
	var userData model.User
	if err := ur.db.Where(model.User{Username: username}).First(&userData).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		} else {
			return false, err
		}
	}
	return userData.ID != "", nil
}
