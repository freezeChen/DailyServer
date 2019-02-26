/*
   @Time : 2019-01-30 16:10
   @Author : frozenchen
   @File : constant
   @Software: DailyServer
*/
package constant

import "time"

const (
	MICRO_IM_SRV    = "vip.frozens.srv.im"
	MICRO_LOGIC_SRV = "vip.frozens.srv.logic"
	MICRO_JOB_SRV   = "vip.frozens.srv.job"
	MICRO_TTL       = 30 * time.Second
	MICRO_Interval  = 20 * time.Second

	JOB_TOPIC_SINGLECHAT = "job_singleChat"
	Job_Topic_AuthReply  = "job_authReply"
)
