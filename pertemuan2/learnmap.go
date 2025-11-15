package pertemuan2

import "fmt"

func LearnMap() {
	person := map[string]string{
		"nama": "andi",
		"umur": "23",
	}

	fmt.Println(person)

	person["alamat"] = "jakarta"

	fmt.Println(person)

	fmt.Println(person["nama"])
}
