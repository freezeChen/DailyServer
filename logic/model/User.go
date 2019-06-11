package model

type User struct {
	Id         int64  `json:"id" xorm:"'id'"`                 //null
	Name       string `json:"name" xorm:"'name'"`             //名称
	Account    string `json:"account" xorm:"'account'"`       //账号
	Password   string `json:"password" xorm:"'password'"`     //密码
	CreateTime string `json:"createTime" xorm:"'createTime'"` //创建时间
}

func (User) TableName() string {
	return "user"
}

