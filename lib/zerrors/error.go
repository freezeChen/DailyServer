/*
   @Time : 2019-06-11 10:48
   @Author : frozenchen
   @File : error
   @Software: DailyServer
*/
package zerrors

import "github.com/micro/go-micro/errors"

func NewMsg(text string) error {
	return errors.New("", text, 200)
}
