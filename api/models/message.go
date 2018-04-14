package models

import (
	"time"
)

type Message struct {
	Fid       string    `xorm:"not null pk VARCHAR(50)"`
	Ffromid   string    `xorm:"VARCHAR(50)"`
	Ftoid     string    `xorm:"VARCHAR(50)"`
	Fmessage  string    `xorm:"VARCHAR(500)"`
	Fsendtime time.Time `xorm:"DATETIME"`
	Fstatus   int       `xorm:"comment('0-已接收 1-未接受') INT(1)"`
}
