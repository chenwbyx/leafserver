package models

import "github.com/jinzhu/gorm"

type Login struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// CheckAuth checks if authentication information exists
func GetLoginByName(username string) (bool, error) {
	var auth Login
	err := db.Select("id").Where(Login{Username: username}).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if auth.ID > 0 {
		return true, nil
	}
	return false, nil
}

func CheckLogin(username, password string) (bool, error) {
	var auth Login
	err := db.Select("id").Where(Login{Username: username, Password: password}).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if auth.ID > 0 {
		return true, nil
	}
	return false, nil
}

// AddTag Add a Login
func AddLogin(username, password string) error {
	login := Login{
		Username:     username,
		Password:     password,
	}
	if err := db.Create(&login).Error; err != nil {
		return err
	}
	return nil
}