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

func (u *User)GetName() ([]byte, error) {
	return USER_DB.Get([]byte(u.Id), nil)
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
