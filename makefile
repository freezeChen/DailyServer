
macBuild:
	cd cmd/im;go build;
	cd cmd/job;go build;
	cd cmd/logic;go build;

linuxBuild:
	cd cmd/im;GOOS=linux go build;
	cd cmd/job;GOOS=linux go build;
	cd cmd/logic;GOOS=linux go build;

upload:linuxBuild
	-ssh root@47.106.137.3 "cd /root/daily;./stop.sh";
	scp cmd/im/im root@47.106.137.3:/root/daily/im/im
	scp cmd/job/job root@47.106.137.3:/root/daily/job/job
	scp cmd/logic/logic root@47.106.137.3:/root/daily/logic/logic
	ssh root@47.106.137.3 "cd /root/daily;./start.sh"

test:macBuild
	-kill -9 $(lsof -t -i:8020)
	-kill -9 ` ps -a |grep logic|awk '{print $1}'`
	-kill -9 ` ps -a |grep job|awk '{print $1}'`
	./cmd/im/im &
	./cmd/job/job &
	./cmd/logic/logic &