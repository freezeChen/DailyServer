/*
   @Time : 2019-06-17 14:58
   @Author : frozenchen
   @File : singleInfo
   @Software: DailyServer
*/
package dao

import "dailyserver/logic/model"

func (d *Dao) InsertSingleInfo(msg *model.SingleInfo) (err error) {

	_, err = d.db.InsertOne(msg)
	return err

}
