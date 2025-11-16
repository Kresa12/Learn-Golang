package latihan2

import "fmt"

type ServicePengguna struct {
	aturan []AturanValidasi
}

func NewServicePengguna(aturan []AturanValidasi) *ServicePengguna {
	return &ServicePengguna{
		aturan,
	}
}

func (s *ServicePengguna) LayananRegristrasi(pengguna *Pengguna) {
	for _, aturan := range s.aturan {
		err := aturan.Cek(pengguna)
		if err != nil {
			fmt.Printf("Validasi GAGAL: %v\n", err)
			return
		}
	}
	fmt.Printf("SUKSES: Pengguna %s berhasil didaftarkan.\n", pengguna.Email)
}
