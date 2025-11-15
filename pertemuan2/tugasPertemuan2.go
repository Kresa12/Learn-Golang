package pertemuan2

import "fmt"

type Todo struct {
	Id     int
	Task   string
	Status bool
}

var todoList = []*Todo{
	{0, "Makan", false},
	{1, "mandi", false},
	{2, "berangkat", false},
}
var id = 3

func AddTodo(task string) {
	newTodo := &Todo{
		Id:     id,
		Task:   task,
		Status: false,
	}
	id++
	todoList = append(todoList, newTodo)
}

func (t *Todo) GetAllTodo() []*Todo {
	if len(todoList) == 0 {
		fmt.Println("Tidak ada data")
		return nil
	}
	return todoList
}

func (t *Todo) GetTodoById(id int) *Todo {
	for _, t := range todoList {
		if id == t.Id {
			return t
		}
	}
	return nil
}

func (t *Todo) Delete(id int) error {
	for i, t := range todoList {
		if id == t.Id {
			todoList = append(todoList[:i], todoList[i+1:]...)
		}
	}

	return fmt.Errorf("id %d tidak ada", id)
}
