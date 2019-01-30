/*
   @Time : 2018/9/3 下午4:56
   @Author : frozenchen
   @File : info
   @Software: DailyServer
*/
package lib

type Info struct {
	Id  int32  `json:"id"`
	Sid int32  `json:"sid"`
	Msg string `json:"msg"`
}
