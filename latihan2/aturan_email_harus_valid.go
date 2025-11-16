package latihan2

import (
	"errors"
	"fmt"
	"strings"
)

type AturanEmailHarusValid struct{}

func (a *AturanEmailHarusValid) Cek(pengguna *Pengguna) error {

	if strings.Contains(pengguna.Email, "@") {
		fmt.Println("email valid, ada karakter @")
		return nil
	} else {
		return errors.New("email tidak valid, tidak ada karakter @ \n")
	}
}
