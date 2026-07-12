package main

type Task struct{
	Taskname string `json:"task"`
	Completed bool `json:"completed"`
}

type TaskList struct{
	Message string `json:"message"`
	Tasks map[int]Task `json:"tasks"`
}

type msgResp struct{
	Message string `json:"message"`
}