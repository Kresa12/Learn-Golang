package latihan2

import (
	"errors"
)

type AturanPasswordMinimal struct {
	panjangMinimal int
}

func (a *AturanPasswordMinimal) Cek(pengguna *Pengguna) error {
	if len(pengguna.Password) < 8 {
		return errors.New("Panjang password minimal 8 karakter! \n")
	}
	return nil
}
