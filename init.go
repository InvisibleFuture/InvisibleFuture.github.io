package main

import (
	"bytes"
	"encoding/binary"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"sync"

	"github.com/syndtr/goleveldb/leveldb"
)

var (
	ACCOUNT_ID_DB *leveldb.DB
	ACCOUNT_PW_DB *leveldb.DB

	USER_NAME_DB    *leveldb.DB
	USER_MARK_DB    *leveldb.DB
	USER_PROJECT_DB *leveldb.DB

	PROJECT_NAME_DB   *leveldb.DB
	PROJECT_MARK_DB   *leveldb.DB
	PROJECT_TAGS_DB   *leveldb.DB // 为项目打TAG?
	PROJECT_TIME_DB   *leveldb.DB
	PROJECT_MASTER_DB *leveldb.DB

	LIST_PROJECT_DB *leveldb.DB // 项目列表 所有

	AUTOID_DB *leveldb.DB // 每次都将ID+1并写入

	AUTOID_USER_CH    chan int64
	AUTOID_PROJECT_CH chan int64

	WEB_TOKEN_MAP sync.Map
	WEB_HTML_MAP  sync.Map
)

func init() {
	// 初始化前检查剩余空间与权限

	var err error

	// ACCOUNT
	ACCOUNT_PW_DB, err = leveldb.OpenFile("../data/account_pw", nil)
	if err != nil {
		panic("ACCOUNT_PW_DB INIT ERROR")
	}

	ACCOUNT_ID_DB, err = leveldb.OpenFile("../data/account_id", nil)
	if err != nil {
		panic("ACCOUNT_ID_DB INIT ERROR")
	}

	// USER
	USER_NAME_DB, err = leveldb.OpenFile("../data/user_name", nil)
	if err != nil {
		panic("USER_NAME_DB INIT ERROR")
	}

	USER_MARK_DB, err = leveldb.OpenFile("../data/user_mark", nil)
	if err != nil {
		panic("USER_MARK_DB INIT ERROR")
	}

	USER_PROJECT_DB, err = leveldb.OpenFile("../data/user_project", nil)
	if err != nil {
		panic("USER_PROJECT_DB INIT ERROR")
	}

	// PROJECT
	PROJECT_NAME_DB, err = leveldb.OpenFile("../data/project_name", nil)
	if err != nil {
		panic("PROJECT_NAME_DB INIT ERROR")
	}

	PROJECT_MARK_DB, err = leveldb.OpenFile("../data/project_mark", nil)
	if err != nil {
		panic("PROJECT_MARK_DB INIT ERROR")
	}

	PROJECT_TAGS_DB, err = leveldb.OpenFile("../data/project_tags", nil)
	if err != nil {
		panic("PROJECT_TAGS_DB INIT ERROR")
	}

	PROJECT_MASTER_DB, err = leveldb.OpenFile("../data/project_master", nil)
	if err != nil {
		panic("PROJECT_MASTER_DB INIT ERROR")
	}

	PROJECT_TIME_DB, err = leveldb.OpenFile("../data/project_time", nil)
	if err != nil {
		panic("PROJECT_TIME_DB INIT ERROR")
	}

	// LIST
	LIST_PROJECT_DB, err = leveldb.OpenFile("../data/list_project", nil)
	if err != nil {
		panic("LIST_PROJECT_DB INIT ERROR")
	}

	AUTOID_DB, err = leveldb.OpenFile("../data/autoid", nil)
	if err != nil {
		panic("AUTOID_DB INIT ERROR")
	}

	// 通道初始化
	AUTOID_USER_CH = make(chan int64)
	AUTOID_PROJECT_CH = make(chan int64)

	// 自增数值独立进程初始化
	go autoid("user", AUTOID_USER_CH)
	go autoid("project", AUTOID_PROJECT_CH)

	// HTML 文件初始化到内存或实现一个监听进程
	file, err := os.Open("./html/index.html")
	if err != nil {
		panic("OPEN FILE INDEX.HTML ERROR")
	}
	html, err := ioutil.ReadAll(file)
	if err != nil {
		panic("LAOD FILE INDEX.HTML ERROR")
	}
	WEB_HTML_MAP.Store("home", html)
	// 监听 HTML 文件修改

	// 设定 HTML 储存文档
	// 设定 CSS 储存文档 LESS ?
	// 设定 JS 储存文档

	// 设定 IMAGE 储存文档

	//USER_NAME_DB.Put([]byte("233"),[]byte("Last"), nil)
	//data, err := db.Get([]byte("key"), nil)
	//err = db.Put([]byte("key"), []byte("value"), nil)
	//err = db.Delete([]byte("key"), nil)
	//defer USER_DB.Close()
}

func rwtoken(id string) string {
	token := randSeq(32)
	WEB_TOKEN_MAP.Store(id, token)
	return token
	// 为 token 设置一个过期时间 计时器
}

func autoid(name string, c chan int64) {
	buf := new(bytes.Buffer)

	data, err := AUTOID_DB.Get([]byte(name), nil)
	if err != nil {
		binary.Write(buf, binary.BigEndian, 0)
		data = buf.Bytes()
		err = AUTOID_DB.Put([]byte(name), data, nil)
		if err != nil {
			panic("AUTOID_DB " + name + " INIT ERROR")
		}
		log.Println("计数器初始化", name)
	}

	var sum int64
	binary.Read(bytes.NewBuffer(data), binary.LittleEndian, &sum)
	log.Println("计数器", name, sum)
	for {
		sum++
		c <- sum
		buf = new(bytes.Buffer)
		binary.Write(buf, binary.LittleEndian, sum)
		err = AUTOID_DB.Put([]byte(name), buf.Bytes(), nil)
		if err != nil {
			panic("AUTOID ++ ERROR")
		}
	}
}

func randSeq(n int) string {
	var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
