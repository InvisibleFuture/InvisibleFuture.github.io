package main

type Task struct {
	Id string
	Time string
}

func (t *Task)Get_Task() Task {
	return Task{
		Id: t.Id,
		Time: "32321212"
	}
}
