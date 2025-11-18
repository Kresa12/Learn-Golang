package function

import (
	"tesGoInIntellij/pertemuan2"
	"testing"
)

func TestAdd(t *testing.T) {
	res := Add(10, 20)
	expect := 30
	if res != expect {
		t.Errorf("hasil %d harusnya %d", res, expect)
	}
}

func TestAddTodoList(t *testing.T) {
	panjangAwal := len(pertemuan2.TodoList)
	newTask := "makan"
	pertemuan2.NewAddTodo(newTask)
	panjangSaatIni := len(pertemuan2.TodoList)
	expect := panjangAwal + 1
	if panjangSaatIni != expect {
		t.Errorf("data tidak bertambah, panjang awal %d, panjang saat ini %d : ", panjangAwal, panjangSaatIni)
	}
}
