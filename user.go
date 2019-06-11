package main

import (
	"fmt"
	"sync"
	"github.com/syndtr/goleveldb/leveldb"
)

var (
	TOKEN_MAP sync.Map
	USER_DB *leveldb.DB
)

func init() {
	USER_DB, err := leveldb.OpenFile("../data/user", nil)
	if err != nil {
		panic("USER_DB INIT ERROR")
	}
	//data, err := db.Get([]byte("key"), nil)
	//err = db.Put([]byte("key"), []byte("value"), nil)
	//err = db.Delete([]byte("key"), nil)
	defer USER_DB.Close()

	TOKEN_MAP.Store("TEAM", "HELLO WORLD")
	if val,ok := TOKEN_MAP.Load("TEAM"); ok {
		fmt.Println(val)
	}
}

type User struct {
	Id string
	Token string
}

func (u *User)Authentication() bool {
	val,ok := TOKEN_MAP.Load(u.Id)
	if ok && val == u.Token {
		return true
	}
	return false
}

func (u *User)SetToken() {
	TOKEN_MAP.Store(u.Id, u.Token)
}


/*
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
*/

