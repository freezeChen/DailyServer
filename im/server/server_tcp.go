/*
   @Time : 2019-05-31 15:23
   @Author : frozenchen
   @File : server_tcp
   @Software: DailyServer
*/
package server

import (
	"dailyserver/im/conf"
	"dailyserver/im/service"
	"net"
)

func InitTCP(svc *service.Service, c *conf.Config) (err error) {
	addr, err := net.ResolveIPAddr("tcp", c.TCPPort)
	if err != nil {
		return
	}

	return nil
}
