package pertemuan3

import "fmt"

type RekeningBank struct {
	Nama  string
	saldo int
}

func (r *RekeningBank) GetSaldo() int {
	return r.saldo
}

func (r *RekeningBank) Deposit(amount int) {
	if amount > 0 {
		r.saldo += amount
		fmt.Printf("anda memasukan deposit sebesar %d, maka jumlah saldo anda sekarang adalah : %d \n", amount, r.saldo)
	} else {
		fmt.Println("deposit yang anda masukkan tidak boleh 0")
	}
}

func (r *RekeningBank) Withdraw(amount int) {
	if amount > 0 && amount < r.saldo {
		r.saldo -= amount
		fmt.Printf("anda berhasil menarik uang sejumlah %d, saldo anda sekarang : %d \n", amount, r.saldo)
	} else {
		fmt.Println("nominal penarikan yang anda masukkan tidak boleh 0")
	}
}
