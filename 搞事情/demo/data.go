
// 静态 HTML CSS LESS JS 是存储在 DB 中
// 软件的生命周期, 启动即激活, 恢复初始只需删库(跑路)
// SERVER 的增减也是动态的, 可热拆卸
// 管理身份可操作增加对象, 对象具有属性
// 操作方法却不可以增加..? 即 MOD
// 关于数据挂钩, 读取 主题 时, 附加的标签 MARK
// 重复动作使用模板(管道), 动态动作使用映射
// 初始的通道是重复动作, 供应


package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
	"encoding/json"
	"encoding/binary"
	"github.com/syndtr/goleveldb/leveldb"
)

var (
	DB map[string]*leveldb.DB       // databases
	CA map[string]sync.Map          // cache
	CH map[string]chan int64        // chan
)

func init() {
	var err error

	DB = make(map[string]*leveldb.DB)
	CA = make(map[string]sync.Map)
	CH = make(map[string]chan int64)

	// init databases
	dbList := []string{"counter", "user", "project"}
	for _, name := range dbList {
		DB[name], err = leveldb.OpenFile(name, nil)
		if err != nil { panic("OpenFile error") }
	}

	// init cache

	// init counter
	chanList := []string{"user", "project"}
	for _, name := range chanList {
		CH[name] = make(chan int64)
		go func(name string, c chan int64){
			var err error
			var sum int64

			buf := make([]byte, 8)

			data, err := DB["counter"].Get([]byte(name), nil)
			if err != nil {
				binary.BigEndian.PutUint64(buf, uint64(sum))
				err = DB["counter"].Put([]byte(name), buf, nil)
				if err != nil { panic("DB Put error") }
			} else {
				sum = int64(binary.BigEndian.Uint64(data))
				log.Println("计数器", name, sum)
			}

			for {
				sum++
				c <- sum
				binary.BigEndian.PutUint64(buf, uint64(sum))
				err = DB["counter"].Put([]byte(name), buf, nil)
				if err!= nil { panic("counter error") }
			}
		}(name, CH[name])
	}

	// init Object
	//obList := []string{"user", "project"}
	//for _, name := range obList {}
}

type Object struct {
	Name string
	ID   string
}

func (o Object) Master() []string {
	//value, err := DB[o.Name].Get([]byte(o.ID), nil)
	//if err != nil {
	//	return nil, err
	//}
	return []string{"323", "12", "666"}
}

type User struct {
	ID string
}

func (u User) Create(o Object, data string) error {
	// 检查 o.Name 是否存在, 从server层传递过来必然要存在 所以不必在这层检查
	// 创建应当此处生成 ID, 并返回给外部
	return DB[o.Name].Put([]byte(o.ID), []byte(data), nil)
}

// Delete Object
func (u User) Delete(o Object) error {
	var err error
	list := o.Master()
	for _, id := range list {
		// is []byte !? no..
		if id == o.ID {
			return DB[o.Name].Delete([]byte(o.ID), nil)
		}
	}
	return err
}

// Rewrite user 只是无权修改其他角色的, 并非没有这个选项
func (u User) Rewrite(o Object, data string) error {
	//if u.ID != o.ID {
	//	return nil
	//}
	return DB[o.Name].Put([]byte(o.ID), []byte(data), nil)
}

// Master : The owner of the user is himself
func (u User) Master() string {
	return u.ID
}

// Signin is user
func (u User) Signin(account, password string) (string, error) {
	//data, err := DB["account"].Get([]byte(account), nil)
	//if err != nil {
	//	return "", err
	//}
	// 分割 data 为 pw 和 salt
	pw := "data"
	if password != pw {
		return "err is pw", nil
	}
	// 设置 token
	// 返回 token true
	return "ok rep token", nil
}

func main() {
	log.Println("start w~")
	http.HandleFunc("/pong", func(w http.ResponseWriter, r *http.Request) {
		num := <-CH["user"]
		log.Println(num)
	})
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		var data interface{}

		// 登录需要的信息, 若是多次操作需要输入验证码
		//[]string{"account", "password", "code"}

		// 注册需要的信息, 限制注册没有意义
		list := []string{"account", "password", "code"}
		for _, key := range list {
			value := r.FormValue(key)
			log.Println(value)
		}

		exe := r.FormValue("exe")
		switch exe {
		case "signin":
			account := r.FormValue("account")
			password := r.FormValue("password")
			// 检测不准为空, 另验证码
			var u User
			token, err := u.Signin(account, password)
			if err != nil {
				data = struct{Code, Uid, Token string}{Uid:"666", Token:token}
			} else {
				data = struct{Code, Info string}{Code:"888", Info:"error"}
			}
		//case "sigout":
		//case "rename":
		//case "create":
		//	account := r.FormValue("account")
		//	password := r.FormValue("password")
		//	log.Println(account, password)
			// 检测不准为空, 另验证码
			// 检测是否被占用
			//var o Object
			//var u User
			//err := u.Create(Object{Name: "UserAccount", ID: newid})
			//err := u.Create(Object{Name: "UserPassword", ID: newid})
		default:
			// 404
		}

		js, err := json.MarshalIndent(data, "", "\t")
		if err != nil { panic("echo json error") }

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write([]byte(js))
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
	time.Sleep(time.Second * 1)

	// Store the token in the database

	// End all database connections
	for name := range DB {
		DB[name].Close()
	}

	log.Println("done.")
}

// 数据, 前端 是底层抽象
// 上层抽象调动底层抽象, 而非一个列表直接运转
// 核心 控制器
// 生 方法? soul
// 物 生的物态 user
// 物 属性数据
// 包含重新编译前端文件吗? 前端是载入到内存的, 所以必然, 之后监听文件更改?
// 已经初始化, 则开始载入基本状态数据
// 状态预备完毕, 开启服务端口
// 流量涌入, 周期状态调整
// 插件是一种理论抽象 而非外部增补, 插件基于基因库, 无需代码

// 生成伪随机字符串, 用于token
//func randSeq(n int) string {
//	var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
//	b := make([]rune, n)
//	for i := range b {
//		b[i] = letters[rand.Intn(len(letters))]
//	}
//	return string(b)
//}
