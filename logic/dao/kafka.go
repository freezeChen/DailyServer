/*
   @Time : 2019-06-17 15:59
   @Author : frozenchen
   @File : kafka
   @Software: DailyServer
*/
package dao

import (
	"dailyserver/logic/model"
	"encoding/json"
	"github.com/micro/go-micro/broker"
)

type kafka struct {
}


func (k *kafka) SendSingleMsg(info *model.SingleInfo) error {
	bytes, err := json.Marshal(info)
	if err != nil {
		return err
	}
	err = broker.Publish("single", &broker.Message{
		Body: bytes,
	})

	return err
}
