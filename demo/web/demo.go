package main

import (
	"log"
	"net/http"
	"../../demo"
)

func main() {
	log.Println("Go")
	log.Println(demo.Texts)
}
func Create(w http.ResponseWriter, r *http.Request) {
	// 创建一个对象, 对象并不需要具有特定事物, 对象是一个标痕
	if r.Method != "POST" {
		w.Write([]byte("必要pot"))
		return
	}
	user := struct{ name, time string }{"name_s","time_s"}
}
func Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Write([]byte("必须post"))
		return
	}
	var u User
	var id, target string
	u.Id    = r.FormValue("uid")
	u.Token = r.FormValue("token")
	id      = r.FormValue("id")
	target  = r.FormValue("target")

	if u.Id == "" || u.Token == "" || target == "" || id == "" {
		w.Write([]byte("非法参数"))
		return
	}
	var object demo.Object
	switch target {
	case "project": object = demo.Project([]byte(id))
	case "user"   : object = demo.User([]byte(id))
	default: return
	}

	if ok := object.Delete(); !ok {
		w.Write([]byte("失败"))
		return
	}
	w.Write([]byte("成功"))
}
type User struct{
	Id    string
	Token string
}
func (u User)Delete(target, id string) bool {
	// 验证身份
	user := demo.User([]byte(u.Id))
	if ok := user.Token([]byte(u.Token)); !ok {
		return false
	}
	// 验证权限

	var object demo.Object
	switch target {
	case "project": object = demo.Project(id)
	case "user"   : object = demo.User(id)
	default       : return false
	}

	return object.Delete()
}
func Create(w, r) {
	// user init
}
func Rw(w, r) {
	//init
}
