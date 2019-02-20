/*
   @Time : 2018/9/3 下午4:56
   @Author : frozenchen
   @File : info
   @Software: DailyServer
*/
package lib

type Info struct {
	Id  int32  `json:"id"`  //发送者id
	Rid int32  `json:"rid"` //接受者id
	Msg string `json:"msg"`
}
