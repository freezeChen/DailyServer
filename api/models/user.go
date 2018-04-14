package models

import (
	"time"
)

type User struct {
	Fid         string    `xorm:"not null pk VARCHAR(50)"`
	Fname       string    `xorm:"VARCHAR(20)"`
	Fheadicon   string    `xorm:"VARCHAR(100)"`
	Faccount    string    `xorm:"VARCHAR(20)"`
	Fpassword   string    `xorm:"VARCHAR(20)"`
	Fsex        int       `xorm:"comment('1-男,0女') INT(1)"`
	Fcreatetime time.Time `xorm:"DATETIME"`
}
