

build:
	cd cmd/im;GOOS=linux go build;
	cd cmd/job;GOOS=linux go build;
	cd cmd/logic;GOOS=linux go build;

upload:build
	scp cmd/im/im root@47.106.137.3:/root/daily/im/im
	scp cmd/job/job root@47.106.137.3:/root/daily/job/job
	scp cmd/logic/logic root@47.106.137.3:/root/daily/logic/logic