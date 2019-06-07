
package main

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
)

var (
	USER_DB *leveldb.DB
)

func init() {
	USER_DB, err := leveldb.OpenFile("data/user", nil)

	if err != nil {
		return
	}
	//data, err := db.Get([]byte("key"), nil)
	//err = db.Put([]byte("key"), []byte("value"), nil)
	//err = db.Delete([]byte("key"), nil)
	defer USER_DB.Close()
}

// 来自外界的 指令 与设备主体是隔离的
// 权限表与数据表是分离的
// 初始化成功才允许继续主进程
// 此处暂时不考虑user表的热自动化

type Object interface{
	Create(string)
	Delete(string)
	Updata(string)
}

type Item struct {
	Id string
	Token string
}

type User struct {
	Id string
	Token string
}

func (u *User)GetName() string {
	return USER_DB.Get([]byte(u.Id))
}

func (u *User)SetName(new_name string) error {
	return	USER_DB.Put([]byte(u.Id), []byte(new_name), nil)
}

func (u *User)Create() string {
	return "ok"
}

func (u *User)Delete(ids []string) string {
	fmt.Println(ids)
	return "ok"
}
