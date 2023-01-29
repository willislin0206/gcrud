package servcie

import (
	orm "data-collector/database"
	. "data-collector/model"

	log "github.com/sirupsen/logrus"
)

//add
func Insert(user *User) (err error) {

	//orm operate user structure and add into database through repository
	result := orm.Conn.Create(&user)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

//list all
func List() (users []User, err error) {
	if err = orm.Conn.Find(&users).Error; err != nil {
		return
	}
	return
}

//list someone
func ListUser(id string) (user User, err error) {
	log.Info(id)
	if err = orm.Conn.Find(&user, "id = ?", id).Error; err != nil {
		return
	}
	return
}

func UpdateUser(id string, user *User) (err error) {
	log.Info(id)

	//var NewUser =User{ Username: user.Username,Password: user.Password }

	var emptyUser User
	if err = orm.Conn.Find(&emptyUser, "id = ?", id).Error; err != nil {
		return
	}

	if err = orm.Conn.Model(&emptyUser).Update(&user).Error; err != nil {
		return
	}

	return
}

func DeleteUser(id string) (user User, err error) {
	log.Info(id)
	if err = orm.Conn.Delete(&user, "id = ?", id).Error; err != nil {
		return
	}
	return
}
