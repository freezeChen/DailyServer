package models

type Frend struct {
	Fid      string `xorm:"not null pk default '' comment('fid') VARCHAR(50)"`
	Fuserid  string `xorm:"default '' VARCHAR(50)"`
	Ffrendid string `xorm:"default '' comment('好友的FuserID') VARCHAR(50)"`
}
