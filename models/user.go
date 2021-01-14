package models

import "awesomeProject/config"

type User struct {
	Id int `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *User) TableName() string {
	return "user"
}

func GetAllUsers(user *[]User) (err error) {
	if err = config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

func CreateUser(user *User) (err error) {
	if err = config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserById(user *User, key string) (err error){
	if err = config.DB.First(&user, key).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *User) (err error) {
	config.DB.Save(user)
	return nil
}

func DeleteUser(user *User, key string) (err error){
	config.DB.Delete(user, key)
	return nil
}