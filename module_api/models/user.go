package models

import (
	"dailyserver2/lib/jsontime"
)

type User struct {
	Id         int64             `xorm:"pk autoincr BIGINT(12)"`
	Name       string            `xorm:"VARCHAR(20)"`
	Account    string            `xorm:"VARCHAR(20)"`
	Password   string            `xorm:"VARCHAR(32)"`
	Createtime jsontime.JsonTime `xorm:"TIMESTAMP(6)"`
}

func GetUserByID(id int64) (*User, error) {
	var user User
	_, err := engine().ID(id).Get(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserList() ([]*User, error) {
	var users []*User
	err := engine().Find(&users)
	return users, err
}
