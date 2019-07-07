package main

import (
	"fmt"
	"time"
	"sync"
	"bytes"
	"strings"
	"encoding/binary"
	"github.com/syndtr/goleveldb/leveldb"
)
var (
	PROJECT_NAME_DB   *leveldb.DB
	PROJECT_MARK_DB   *leveldb.DB
	PROJECT_TAGS_DB   *leveldb.DB // 为项目打TAG?
	PROJECT_TIME_DB   *leveldb.DB
	PROJECT_MASTER_DB *leveldb.DB

	USER_NAME_DB      *leveldb.DB
	USER_MARK_DB      *leveldb.DB
	USER_PROJECT_DB   *leveldb.DB

	LIST_PROJECT_DB   *leveldb.DB // 项目列表 所有

	AUTOID_DB         *leveldb.DB // 每次都将ID+1并写入

	AUTOID_USER_CH    chan int64
	AUTOID_PROJECT_CH chan int64

	TOKEN_MAP         sync.Map
)
func init() {
	// 初始化前检查剩余空间与权限

	var err error

	// USER
	USER_NAME_DB, err = leveldb.OpenFile("../data/user_name", nil)
	if err != nil { panic("USER_NAME_DB INIT ERROR") }

	USER_MARK_DB, err = leveldb.OpenFile("../data/user_mark", nil)
	if err != nil { panic("USER_MARK_DB INIT ERROR") }

	USER_PROJECT_DB, err = leveldb.OpenFile("../data/user_project", nil)
	if err != nil { panic("USER_PROJECT_DB INIT ERROR") }

	// PROJECT
	PROJECT_NAME_DB, err = leveldb.OpenFile("../data/project_name", nil)
	if err != nil { panic("PROJECT_NAME_DB INIT ERROR") }

	PROJECT_MARK_DB, err = leveldb.OpenFile("../data/project_mark", nil)
	if err != nil { panic("PROJECT_MARK_DB INIT ERROR") }

	PROJECT_TAGS_DB, err = leveldb.OpenFile("../data/project_tags", nil)
	if err != nil { panic("PROJECT_TAGS_DB INIT ERROR") }

	PROJECT_MASTER_DB, err = leveldb.OpenFile("../data/project_master", nil)
	if err != nil { panic("PROJECT_MASTER_DB INIT ERROR") }

	PROJECT_TIME_DB, err = leveldb.OpenFile("../data/project_time", nil)
	if err != nil { panic("PROJECT_TIME_DB INIT ERROR") }

	// LIST
	LIST_PROJECT_DB, err = leveldb.OpenFile("../data/list_project", nil)
	if err != nil { panic("LIST_PROJECT_DB INIT ERROR") }

	AUTOID_DB, err = leveldb.OpenFile("../data/autoid", nil)
	if err != nil { panic("AUTOID_DB INIT ERROR") }

	// 通道初始化
	AUTOID_USER_CH = make(chan int64)
	AUTOID_PROJECT_CH = make(chan int64)

	// 自增数值独立进程初始化
	go autoid("user", AUTOID_USER_CH)
	go autoid("project", AUTOID_PROJECT_CH)

	//USER_NAME_DB.Put([]byte("233"),[]byte("Last"), nil)
	//data, err := db.Get([]byte("key"), nil)
	//err = db.Put([]byte("key"), []byte("value"), nil)
	//err = db.Delete([]byte("key"), nil)
	//defer USER_DB.Close()
}

func autoid(name string, c chan int64) {
	buf := new(bytes.Buffer)

	data, err := AUTOID_DB.Get([]byte(name), nil)
	if err != nil {
		binary.Write(buf, binary.BigEndian, 0)
		data = buf.Bytes()
		err = AUTOID_DB.Put([]byte(name), data, nil)
		if err != nil { panic("AUTOID_DB " + name + " INIT ERROR") }
		fmt.Println("计数器初始化", name)
	}

	var sum int64
	binary.Read(bytes.NewBuffer(data), binary.LittleEndian, &sum)
	fmt.Println("计数器", name, sum)
	for {
		sum++
		c <- sum
		buf = new(bytes.Buffer)
		binary.Write(buf, binary.LittleEndian, sum)
		err = AUTOID_DB.Put([]byte(name), buf.Bytes(), nil)
		if err != nil { panic("AUTOID ++ ERROR") }
	}
}

type Project struct {
	Id      string   `json:"id"`
	Name    string   `json:"name"`
	Time    int64    `json:"time"`
	Item    string   `json:"item"`
	Task    []string `json:"task"`
	Mark    []string `json:"mark"`
	Master  []string `json:"master"`
	Partner []string `json:"partner"`
}

func (p Project)Create(name,master string) {
	//id := <-AUTOID_PROJECT_CH
	err := PROJECT_NAME_DB.Put([]byte(p.Id), []byte(name), nil)
	if err != nil { panic("PROJECT_NAME_DB CREATE ERROR") }

	t := make([]byte, 8)
	binary.BigEndian.PutUint64(t, uint64(time.Now().Unix()))
	err = PROJECT_TIME_DB.Put([]byte(p.Id), []byte(t), nil)
	if err != nil { panic("PROJECT_TIME_DB CREATE ERROR") }

	err = PROJECT_MASTER_DB.Put([]byte(p.Id), []byte(master), nil)
	if err != nil { panic("PROJECT_MASTER_DB CREATE ERROR") }

	//err = PROJECT_TIME_DB.Put([]byte(p.Id), []byte(time), nil)
	//if err != nil { panic("PROJECT_TIME_DB CREATE ERROR") }
}
type ListProject struct{
	Id   string `json:"id"`
	Name string `json:"name"`
	Item string `json:"item"`
	User string `json:"user"`
}
func (l List)GetProjects() ([]ListProject, error) {
	LIST_PROJECT_DB.Put([]byte(l.Id), []byte("DEMO_key DEMO_BEY"), nil)
	data, err := LIST_PROJECT_DB.Get([]byte(l.Id), nil)
	ids := strings.Fields(string(data))
	p := []ListProject{}
	for _, id := range ids {
		PROJECT_NAME_DB.Put([]byte(id), []byte("DEMO_DATA"), nil)
		data, err = PROJECT_NAME_DB.Get([]byte(id), nil)
		if err != nil { panic("PROJECT_NAME_DB GET ERROR") }
		p = append(p, ListProject{Id:id, Name: string(data)})
	}
	return p, err
}
type List struct {
	Id string
}
func (l List)GetProject() ([]byte, error) {
	LIST_PROJECT_DB.Put([]byte(l.Id), []byte("VVVALUE"), nil)
	return LIST_PROJECT_DB.Get([]byte(l.Id), nil)
}
//func (p Project)Get() {
//	data, err := PROJECT_NAME_DB.Get([]byte(p.Id), nil)
//	if err != nil { panic("404") }
//	data, err := PROJECT_INFO_DB.Get([]byte(p.Id), nil)
//	if err != nil { panic("404") }
//	data, err := PROJECT_POST_DB.Get([]byte(p.Id), nil)
//	//attach reply.. !! 每当有回复时计数器++
//}
func (p Project)GetName()([]byte, error){
	PROJECT_NAME_DB.Put([]byte(p.Id), []byte("VVVNAME"), nil)
	return PROJECT_NAME_DB.Get([]byte(p.Id), nil)
}
func (p Project)GetMark()([]byte, error){
	return PROJECT_MARK_DB.Get([]byte(p.Id), nil)
}
func (p Project)SetName(){}
func (p Project)SetMark(){}
func (p Project)Delete() error {
	id := []byte(p.Id)
	if err := PROJECT_NAME_DB.Delete(id, nil); err != nil {
		return err
	}
	if err := PROJECT_MARK_DB.Delete(id, nil); err != nil {
		return err
	}
	//if err := PROJECT_ID_DB.Delete(id, nil); err != nil {
	//	return err
	//}
	return nil
	// 删评论 由于已经验证了对主贴的权限, 不必管跟贴权限
}
