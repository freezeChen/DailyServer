/*
    @Time : 2018/9/3 下午4:56 
    @Author : frozenchen
    @File : info
    @Software: dailyserver2
*/
package lib

type Info struct {
	Id  int64  `json:"id"`
	Sid int64  `json:"sid"`
	Msg []byte `json:"msg"`
}
