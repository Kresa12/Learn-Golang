package latihan2

type AturanValidasi interface {
	Cek(pengguna *Pengguna) error
}
