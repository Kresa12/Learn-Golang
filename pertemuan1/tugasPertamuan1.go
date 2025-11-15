package pertemuan1

import (
	"errors"
	"fmt"
)

func Tambah(angkaA, angkaB int) int {
	return angkaA + angkaB
}

func Kurang(angkaA, angkaB int) int {
	return angkaA - angkaB
}

func Kali(angkaA, angkaB int) int {
	return angkaA * angkaB
}

func Bagi(angkaA, angkaB int) int {
	return angkaA / angkaB
}

func ErrorHandilngfunc(input int) (string, error) {
	if input < 0 {
		return "angka yang di input tidak boleh kurang dari 0", errors.Join()
	}

	return fmt.Sprintf("angka yang kamu masukan : % d", input), nil
}
