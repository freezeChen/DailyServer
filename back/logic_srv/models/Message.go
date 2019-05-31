package models

import "github.com/micro/go-micro/errors"

type Message struct {
	Id     int64  `json:"id" xorm:"pk 'id'"`      //null
	Uid    int32  `json:"uid" xorm:"'uid'"`       //发送者id
	Rid    int32  `json:"rid" xorm:"'rid'"`       //接受者id
	Type   int32  `json:"type" xorm:"'type'"`     //消息类型(1:单聊信息)
	Msg    string `json:"msg" xorm:"'msg'"`       //信息
	Status int32  `json:"status" xorm:"'status'"` //状态(0:未发送,1:已发送)
}

func (Message) TableName() string {
	return "message"
}

func InsertMsg(msg *Message) error {
	insertOne, err := Engine().InsertOne(msg)
	if err != nil {
		return err
	}

	if insertOne != 1 {
		return errors.New("", "add msg error", 200)

	}
	return nil
}
