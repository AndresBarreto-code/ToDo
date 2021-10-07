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

func (tasksList *TaskList) AddTask(task *Task) {
	tasksList.tasks = append(tasksList.tasks, task)
}

func (tasksList *TaskList) RemoveTask(index int) {
	tasksList.tasks = append(tasksList.tasks[:index], tasksList.tasks[index+1:]...)
}

func (tasksList *TaskList) Print() {
	for _, task := range tasksList.tasks {
		fmt.Println("Name: ", task.name)
		fmt.Println("Description: ", task.description)
		fmt.Println("Completed?: ", task.complete)
	}
}

func (tasksList *TaskList) PrintComplete() {
	for _, task := range tasksList.tasks {
		if task.complete {
			fmt.Println("Name: ", task.name)
			fmt.Println("Description: ", task.description)
		}
	}
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
	task0 := NewTask("Finish go course", "Finish my go course from Platzi")
	task1 := NewTask("Finish Python course", "Finish my Python course from Platzi")
	task2 := NewTask("Finish Javascript course", "Finish my Javascript course from Platzi")

	fmt.Printf("%+v\n", task0)
	task0.UpdateName("Finish go course!!!")
	task0.UpdateDescription("Finish my go course from Platzi!!!")
	task0.Complete()
	fmt.Printf("%+v\n", task0)

	taskList := NewTaskList("Andres")
	fmt.Printf("%+v\n", taskList)
	taskList.AddTask(task0)
	taskList.AddTask(task1)
	taskList.AddTask(task2)
	fmt.Printf("%+v\n", taskList)
	fmt.Println(taskList.tasks)
	fmt.Println(taskList.tasks[0])
	taskList.RemoveTask(1)
	fmt.Println(taskList.tasks)
	taskList.AddTask(task1)
	fmt.Println(taskList.tasks)
	taskList.Print()
	taskList.PrintComplete()

}
