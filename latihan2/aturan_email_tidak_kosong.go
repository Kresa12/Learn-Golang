package latihan2

import (
	"errors"
)

type AturanEmailTidakKosong struct{}

func (a *AturanEmailTidakKosong) Cek(pengguna *Pengguna) error {
	if len(pengguna.Email) == 0 {
		return errors.New("Email tidak boleh kosong! \n")
	}
	return nil
}
