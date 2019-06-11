package main

import (
	"github.com/syndtr/goleveldb/leveldb"
)

var (
	PROJECT_NAME_DB *leveldb.DB
	PROJECT_MARK_DB *leveldb.DB
)

type Project struct {
	Id      string   `json:"id"`
	Name    string   `json:"name"`
	Time    string   `json:"time"`
	Item    string   `json:"item"`
	Task    []string `json:"task"`
	Mark    []string `json:"mark"`
	Master  []string `json:"master"`
	Partner []string `json:"partner"`
}

func init() {
	var err error
	PROJECT_NAME_DB, err = leveldb.OpenFile("../data/peoject_name", nil)
	if err != nil {
		panic("PROJECT_NAME_DB INIT ERROR")
	}
	PROJECT_MARK_DB, err = leveldb.OpenFile("../data/peoject_mark", nil)
	if err != nil {
		panic("PROJECT_NAME_DB INIT ERROR")
	}
}
func (p *Project)GetName()([]byte, error){
	return PROJECT_NAME_DB.Get([]byte(p.Id), nil)
}
func (p *Project)GetMark()([]byte, error){
	return PROJECT_MARK_DB.Get([]byte(p.Id), nil)
}

func (p *Project)SetName(){}
func (p *Project)SetMark(){}


