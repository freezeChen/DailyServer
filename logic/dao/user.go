/*
   @Time : 2019-06-11 10:42
   @Author : frozenchen
   @File : user
   @Software: DailyServer
*/
package dao

import "dailyserver/logic/model"

func (d *Dao) Login(account, password string) *model.User {
	var user model.User
	user.Account = account
	user.Password = password

	has, err := d.db.Get(&user)
	if err != nil {
		return nil
	}
	if !has {
		return nil
	}
	return &user
}

func (d *Dao) GetUser(id int64) *model.User {
	var user model.User
	has, err := d.db.ID(id).Get(&user)
	if err != nil {
		return nil
	}
	if !has {
		return nil
	}
	return &user
}
