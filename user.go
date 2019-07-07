package main

import (
	"fmt"
	"bytes"
	"strings"
)

type User struct {
	Id string
	Token string
}

func (u User)Create(target, id string) {
	//id := <-AUTOID_USER_CH
	//USER_
}

// 角色可操作的项目 || 可被操作的项目
func (u User)Delete(object interface{}) bool {
	// 验证身份
	if t, ok := TOKEN_MAP.Load(u.Id); !ok || t != u.Token {
		return false
	}

	// 验证权限
	if ok := object.Master(u.Id); !ok {
		return false
	}

	// 操作删除
	if err := object.Delete(); err != nil {
		return false
	}

	return true
}
func (u User)Master(id string) ([]string, bool) {
	data, err := USER_MASTER_DB.Get([]byte(u.Id), nil)
	if err != nil { return false }
	list := strings.Fields(string(data))
	for _, uid := range list {
		if id == uid {
			return true
		}
	}
	return false
}

func (u User)ProjectPush(id string) {
	// 此处也可以考虑直接对 []byte 拼接,而不是转回 string
	data, err := USER_PROJECT_DB.Get([]byte(u.Id), nil)
	if err != nil { panic("USER_PROJECT_DB PULL ERROR") }
	//list := strings.Fields(string(data))
	var buf bytes.Buffer
	buf.WriteString(string(data))
	buf.WriteString(" ")
	buf.WriteString(id)
	err = USER_PROJECT_DB.Put([]byte(u.Id), []byte(buf.String()), nil)
	if err != nil { panic("USER_PROJECT_DB PUSH ERROR") }
}

func (u *User)Authentication() bool {
	val,ok := TOKEN_MAP.Load(u.Id)
	fmt.Println(val," ",u.Token)
	if ok && val == u.Token {
		return true
	}
	return false
}

func (u *User)SetToken() {
	TOKEN_MAP.Store(u.Id, u.Token)
	// 注意还需要重置有效时间
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

