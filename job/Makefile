build:
	go build -o ./cmd/cmd ./cmd/main.go

genproto:
		cd proto;protoc  --micro_out=. -I . -I $(GOPATH)/src --go_out=. *.proto;
