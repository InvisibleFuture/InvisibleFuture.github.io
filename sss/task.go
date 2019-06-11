package main

import (
	"github.com/syndtr/goleveldb/leveldb"
)

var (
	TASK_DB *leveldb.DB
)

func init() {
	TASK_DB, err := leveldb.OpenFile("data/task", nil)
	if err != nil {
		panic("TASK_DB INIT ERROR")
	}
	defer TASK_DB.Close()
}

type Task struct {
	Id string
	Time string
	say []say
}

func Get_Task(id string) Task {
	return Task{
		Id: id,
		Time: "32321212"
	}
}

func (t *Task)SetTime(new_time string) error {
	return TASK_DB.Put([]byte(t.Id), []byte(new_time), nil)
}

/*
task

*/
