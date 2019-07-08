package main

import (
	"bytes"
)


type User []byte

func (u User)Create(target, id string) {
	//id := <-AUTOID_USER_CH
	//USER_
}

func (u User)Delete(id []byte) bool {
	// 验证权限
	// 操作删除
	// 回执
	return true
}

/** 这是其他项的权限判断方法
func (u User)Master(id string) bool {
	data, err := USER_MASTER_DB.Get([]byte(u.Id), nil)
	if err != nil { return false }
	list := strings.Fields(string(data))
	for _, uid := range list {
		if id == uid {
			return true
		}
	}
	return false
}**/

func (u User)ProjectPush(id string) {
	data, err := USER_PROJECT_DB.Get(u, nil)
	if err != nil { panic("USER_PROJECT_DB PULL ERROR") }
	//list := strings.Fields(string(data))
	var buf bytes.Buffer
	buf.WriteString(string(data))
	buf.WriteString(" ")
	buf.WriteString(id)
	err = USER_PROJECT_DB.Put(u, []byte(buf.String()), nil)
	if err != nil { panic("USER_PROJECT_DB PUSH ERROR") }
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
/**
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
**/
