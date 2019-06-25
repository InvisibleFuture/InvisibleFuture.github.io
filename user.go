package main

import (
	"fmt"
	"sync"
	"strings"
	"github.com/syndtr/goleveldb/leveldb"
)

var (
	TOKEN_MAP sync.Map
	USER_NAME_DB *leveldb.DB

	USER_MARK_DB *leveldb.DB
	USER_PROJECT_DB *leveldb.DB
)

func init() {
	var err error
	USER_NAME_DB, err = leveldb.OpenFile("../data/user_name", nil)
	if err != nil { panic("USER_NAME_DB INIT ERROR") }

	USER_MARK_DB, err = leveldb.OpenFile("../data/user_mark", nil)
	if err != nil { panic("USER_MARK_DB INIT ERROR") }

	USER_PROJECT_DB, err = leveldb.OpenFile("../data/user_project", nil)
	if err != nil { panic("USER_PROJECT_DB INIT ERROR") }

	USER_NAME_DB.Put([]byte("233"),[]byte("Last"), nil)

	//data, err := db.Get([]byte("key"), nil)
	//err = db.Put([]byte("key"), []byte("value"), nil)
	//err = db.Delete([]byte("key"), nil)
	//defer USER_DB.Close()

	TOKEN_MAP.Store("TEAM", "HELLO WORLD")
	if val,ok := TOKEN_MAP.Load("TEAM"); ok {
		fmt.Println(val)
	}
}

type User struct {
	Id string
	Token string
}

func (u User)Create() {
	//id := <-AUTOID_USER_CH
	//USER_
}

func (u *User)AddProject(id []string) {
	//data, err := USER_PROJECT_DB.Get([]byte(u.Id), nil)
	//if err != nil { return nil, err }
	//list := strings.Fields(string(data))
	// 不必拆分直接追加? 需要对比重复的存在
	//fmt.Println(list) //[]string
	//对比是否有重复
	//向list 追加两个string
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


//func (u *User)GetProject() ([]string, error) {
//	project, err := USER_PROJECT_DB.Get([]byte(u.Id), nil)
//	if err != nil { return nil, err }
//	list := strings.Fields("hello widuu golang")
//	fmt.Println(list, project)
//	return list, err
//}

//func (u *User)SetProject() ([]string, error) {
//	var list []string
//	data, err := USER_PROJECT_DB.Get([]byte(u.Id), nil)
//	if err != nil { return list, err }
//	list = strings.Fields("hello widuu golang")
//	fmt.Println(list, data)
//	return list, err
//}

func (u *User)GetMark() ([]string, error) {
	mark, err := USER_MARK_DB.Get([]byte(u.Id), nil)
	if err != nil {
		return nil, err
	}
	list := strings.Fields("hello widuu golang")
	fmt.Println(list, mark)
	return list, err
}

func (u *User)SetMark() {
}

func (u *User)GetUser() ([]byte, error) {
	return USER_NAME_DB.Get([]byte(u.Id), nil)
}

func (u *User)GetProject() ([]string, error) {
	var list []string
	data, err := USER_PROJECT_DB.Get([]byte(u.Id), nil)
	if err != nil { return list, err }
	list = strings.Fields(string(data))
	fmt.Println(list)
	return list, err
}

