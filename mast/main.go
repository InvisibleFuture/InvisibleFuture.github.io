package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
 * 一个标准对象具有众多未知或已知属性
 * 一个标准对象具有众多未知或已知方法
 * 创建一篇文章 *
 * 创建一个角色 *
 * 角色创建一篇文章
 * 角色创建一个角色
 * 角色修改一个角色, 通常只有权修改自己
 * 文章创建一个角色, ???
 * 文章创建一个文章, 当然除非文章是活体
 * 活体的定义是具有意志的, 交互数据的
 * 文章和角色是不同的结构体, 具有不同的方法函数
**/

// Object 是个虚拟对象, 即接口
type Object interface {
	Delete() bool
	Master() []string
}

// User 是用户身份的抽象, 通常仅需要ID作为标识
type User struct {
	ID    string
	Token string
}

// Project ihts aaa
type Project struct {
	ID string
}

// Delete is Project 对象删除自身的方法
func (p Project) Delete() bool {
	// 删除相关联的数据, 如其他角色的收藏?
	// 对目标删除
	return true
}

// Master list is Project
func (p Project) Master() []string {
	return []string{"233", "3234", "423"}
}

// Delete User
func (u User) Delete(o Object) bool {
	// 检查是否具有删除权限
	list := o.Master()
	for _, id := range list {
		if id == u.ID {
			return o.Delete()
		}
	}
	return false
}

func main() {
	http.HandleFunc("/project", func(w http.ResponseWriter, r *http.Request) {
		// 判断是否 POST // 访问所有对象 而不只是 porject
		user := User{
			ID:    r.FormValue("uid"),
			Token: r.FormValue("token"),
		}
		project := Project{
			ID: r.FormValue("pid"),
		}
		message := false
		execute := r.FormValue("exe")
		switch execute {
		case "delete":
			message = user.Delete(project)
		//case "create": message = user.Create(Project{id:r.FormValue})
		default:
			message = false
		}
		// 操作只有成败, 无需理由
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		//w.WriteHeader(code)
		html := "err"
		if message {
			html = "ok"
		}
		w.Write([]byte(html))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 使用 SQL 时, 如有注入风险的情况, 应在与SQL接口前对数据安全转换或过滤
		// 所有操作应使用 POST
		// 所有读取应使用 GET
		// 常规 POST 应携带 TOKEN
		// 不使用 COOKIES
		// WEB 层没有对象 结构体
		// 访问所有对象, 而不只是 project
		// API 访问并不是实体对象, 不应占用实体对象的 URL
		html := "index !"
		op   := r.FormValue("op")
		obj  := r.FormValue("obj")
		switch op {
		case "delete":
			ids := r.FormValue("id")
			user.Delete(obj, ids)
		case "create":
			// 创建对象时, 对象的组成是不同的, 如何..
			// 增加tag时, tag的对象是不同的: 但tag属于对象 不是级联关系
			// 动态创建新的仓库, 而不是写成常量
			// 禁止使用常量!
			user.Create(obj, "Information Flow")
		case "mark":
			ids := r.FormValue("id")
			user.Mark(obj, ids)
		default:
			w.WriteHeader(404)
			html = "<h1>404</h1>"
		}

		// URL必要是地址, 而非操作, 参数才是操作/ 与函数排列概念是错位的
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		//w.WriteHeader(code)
		w.Write([]byte(html))
	})

	s := &http.Server{
		Addr:           ":80",
		Handler:        http.DefaultServeMux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		log.Println(s.ListenAndServe())
		log.Println("server shutdown")
	}()

	// Handle SIGINT and SIGTERM.
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)

	// Stop the service gracefully.
	log.Println(s.Shutdown(nil))

	// Wait gorotine print shutdown message
	// time.Sleep(time.Second * 1)
	// 结束所有数据库连接
	log.Println("done.")
}
