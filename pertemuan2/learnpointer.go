package pertemuan2

import "fmt"

func LearnPointer() {

	a := 10
	var p *int = &a

	*p = 20
	fmt.Println(a)
	fmt.Println(*p)
}
