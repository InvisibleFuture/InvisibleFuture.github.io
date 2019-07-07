package main

import (
	"time"
	"strings"
	"encoding/binary"
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
