package models

type User struct {
	Id       int64  `json:"id" xorm:"pk 'id'"`          //null
	Account  string `json:"account" xorm:"'account'"`   //null
	Password string `json:"password" xorm:"'password'"` //null
	Nickname string `json:"nickname" xorm:"'nickname'"` //null
	Sex      int64  `json:"sex" xorm:"'sex'"`           //null
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
