/*
   @Time : 2018/8/29 下午6:50
   @Author :
   @File : user
   @Software: DailyServer
*/
package cache

import (
	"DailyServer/back/commons/gredis"
	"DailyServer/back/commons/util"
)

const Online_key = "online:"

func SaveOnlineUser(id int32) error {
	return gredis.Set(Online_key+util.ToString(id), "1", 30*60)
}

func UserIsOnline(id int32) bool {
	return gredis.Exists(Online_key + util.ToString(id))
}
