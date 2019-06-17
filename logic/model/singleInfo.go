/*
   @Time : 2019-06-17 14:57
   @Author : frozenchen
   @File : singleInfo
   @Software: DailyServer
*/
package model

type SingleInfo struct {
	Id   int64 `xorm:"pk"`
	Uid  int64
	ToId int64
	Msg  string
	Time string
}
