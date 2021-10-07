package main

import "fmt"

type Task struct {
	name        string
	description string
	complete    bool
}

type TaskList struct {
	tasks []*Task
	owner string
}

func (taksList *TaskList) AddTask(task *Task) {
	taksList.tasks = append(taksList.tasks, task)
}

func (t *Task) Complete() {
	t.complete = true
}

func (t *Task) UpdateDescription(description string) {
	t.description = description
}

func (t *Task) UpdateName(name string) {
	t.name = name
}

func NewTask(name string, description string) *Task {
	return &Task{
		name:        name,
		description: description,
		complete:    false,
	}
}

func NewTaskList(owner string) *TaskList {
	return &TaskList{
		tasks: []*Task{},
		owner: owner,
	}
}
func main() {
	task := NewTask("Finish go course", "Finish my go course from Platzi")

	fmt.Printf("%+v\n", task)
	task.UpdateName("Finish go course!!!")
	fmt.Printf("%+v\n", task)
	task.UpdateDescription("Finish my go course from Platzi!!!")
	fmt.Printf("%+v\n", task)
	task.Complete()
	fmt.Printf("%+v\n", task)

	taskList := NewTaskList("Andres")
	fmt.Printf("%+v\n", taskList)
	taskList.AddTask(task)
	fmt.Printf("%+v\n", taskList)
	fmt.Println(taskList.tasks)
	fmt.Println(taskList.tasks[0])

}
