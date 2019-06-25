package main

import (
	"fmt"
	"bytes"
	"encoding/binary"
	"github.com/syndtr/goleveldb/leveldb"
)

var (
	PROJECT_NAME_DB *leveldb.DB
	PROJECT_MARK_DB *leveldb.DB
	PROJECT_LIST_DB *leveldb.DB // 项目列表 所有
	PROJECT_TAGS_DB *leveldb.DB // 为项目打TAG?

	AUTOID_DB *leveldb.DB // 每次都将ID+1并写入

	AUTOID_USER_CH chan int64
	AUTOID_PROJECT_CH chan int64
)

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

func init() {
	// 初始化前检查剩余空间与权限

	var err error
	PROJECT_NAME_DB, err = leveldb.OpenFile("../data/project_name", nil)
	if err != nil { panic("PROJECT_NAME_DB INIT ERROR") }

	PROJECT_MARK_DB, err = leveldb.OpenFile("../data/project_mark", nil)
	if err != nil { panic("PROJECT_MARK_DB INIT ERROR") }

	PROJECT_TAGS_DB, err = leveldb.OpenFile("../data/project_tags", nil)
	if err != nil { panic("PROJECT_TAGS_DB INIT ERROR") }

	PROJECT_LIST_DB, err = leveldb.OpenFile("../data/project_list", nil)
	if err != nil { panic("PROJECT_LIST_DB INIT ERROR") }

	AUTOID_DB, err = leveldb.OpenFile("../data/autoid", nil)
	if err != nil { panic("AUTOID_DB INIT ERROR") }

	// 创建通道
	AUTOID_USER_CH = make(chan int64)
	AUTOID_PROJECT_CH = make(chan int64)

	// AUTOID INIT GO
	go autoid("user", AUTOID_USER_CH)
	go autoid("project", AUTOID_PROJECT_CH)

	var sum int64
	sum = <-AUTOID_USER_CH
	fmt.Println("get", sum)

	sum = <-AUTOID_USER_CH
	fmt.Println("get", sum)

	sum = <-AUTOID_USER_CH
	fmt.Println("get", sum)

	sum = <-AUTOID_USER_CH
	fmt.Println("get", sum)
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

//func (p Project)Delete() error {
	// 包含的
	//Mark project的存在, 如何得知新增了哪些呢? 如何新增呢..
	//delete("project", "id")
	//get project_mark // ["32","33","34"]
	// 判断是否删除时的权限表呢,,  也是一种内在关联,, 当角色要删除一个对象, 不必读取对象本身, 而去权限表验证, 前端角色又是怎么知道自己有没有权限呢
	// 关联的
	//USER_PROJECT_DB.Delete([]byte(p.Id), nil)
//}

func (p *Project)GetName()([]byte, error){
	return PROJECT_NAME_DB.Get([]byte(p.Id), nil)
}
func (p *Project)GetMark()([]byte, error){
	return PROJECT_MARK_DB.Get([]byte(p.Id), nil)
}

func (p *Project)SetName(){}
func (p *Project)SetMark(){}

func (p *Project)Delete() error {
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
