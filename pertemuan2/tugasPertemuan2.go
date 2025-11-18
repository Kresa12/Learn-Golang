package pertemuan2

import "fmt"

type Todo struct {
	Id     int
	Task   string
	Status bool
}

var TodoList = []*Todo{
	{0, "Makan", false},
	{1, "mandi", false},
	{2, "berangkat", false},
}
var id = 3

func NewAddTodo(task string) *Todo {
	newTodo := &Todo{
		Id:     id,
		Task:   task,
		Status: false,
	}
	id++
	TodoList = append(TodoList, newTodo)
	return newTodo
}

func (t *Todo) GetAllTodo() []*Todo {
	if len(TodoList) == 0 {
		fmt.Println("Tidak ada data")
		return nil
	}
	return TodoList
}

func (t *Todo) GetTodoById(id int) *Todo {
	for _, t := range TodoList {
		if id == t.Id {
			return t
		}
	}
	return nil
}

func (t *Todo) Delete(id int) error {
	for i, t := range TodoList {
		if id == t.Id {
			TodoList = append(TodoList[:i], TodoList[i+1:]...)
		}
	}

	return fmt.Errorf("id %d tidak ada", id)
}
