package models

type Info struct {
	FUserID string
	FInfo   ChatInfo
}

type ChatInfo struct {
	FTOID    string
	FUser    *User
	FType    int
	FContent string
	FTime    string
}
