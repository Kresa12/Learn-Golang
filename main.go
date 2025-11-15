package main

import (
	"fmt"
	"tesGoInIntellij/latihan"
)

type Eksporter interface {
	Ekspor(data string) error
}

type FileEksporter struct {
	namaFile string
}

type CsvEksporter struct {
	FileEksporter
}

func (c *CsvEksporter) Ekspor(data string) error {
	fmt.Printf("Mengekspor data %s ke file CSV : %s \n ", data, c.namaFile)
	return nil
}

type JsonEksporter struct {
	FileEksporter
}

func (c *JsonEksporter) Ekspor(data string) error {
	fmt.Printf("Mengekspor data %s ke file Json : %s \n ", data, c.namaFile)
	return nil
}

type ConsoleEksporter struct {
}

func (c *ConsoleEksporter) Ekspor(data string) error {
	fmt.Printf("Mengekspor data ke CONCOLE %s \n ", data)
	return nil
}

func jalankanEksport(dataLaporan string, eksportir Eksporter) {
	err := eksportir.Ekspor(dataLaporan)
	if err != nil {
		return
	}
}

func main() {

	order1 := latihan.Order{Id: "1", Amount: 300, EmailUser: "bagas@gmail.com"}
	procCreditCard := &latihan.CreditCardProcessor{}
	notifEmail := &latihan.EmailNotifier{}
	layananOrder1 := latihan.NewOrderService(procCreditCard, notifEmail)
	layananOrder1.OrderProces(&order1)

	order2 := latihan.Order{Id: "2", Amount: 200, EmailUser: "fikal@gmail.com"}
	procbank := &latihan.BankTransferProcessor{}
	notifEmail = &latihan.EmailNotifier{}
	layananOrder2 := latihan.NewOrderService(procbank, notifEmail)
	layananOrder2.OrderProces(&order2)

	order3 := latihan.Order{Id: "3", Amount: 150, EmailUser: "panji@gmail.com"}
	procbank = &latihan.BankTransferProcessor{}
	notifWhatsApp := &latihan.WhatsAppNotifier{}
	layananOrder3 := latihan.NewOrderService(procbank, notifWhatsApp)
	layananOrder3.OrderProces(&order3)

	//db := pengambilandata.NewDatabaseStore()
	//db.Data["1"] = "Andi"
	//userServicedb := pengambilandata.NewUserService(db)
	//userServicedb.ShowUserName("1")
	//userServicedb.ShowUserName("2")
	//
	//cache := pengambilandata.NewCacheStore()
	//cache.Data["2"] = "Fikal"
	//userServicecache := pengambilandata.NewUserService(cache)
	//userServicecache.ShowUserName("2")
	//userServicecache.ShowUserName("3")
	//
	//api := pengambilandata.ApiStore{}
	//userServiceapi := pengambilandata.NewUserService(&api)
	//userServiceapi.ShowUserName("9")

	//eksportirCsv := CsvEksporter{FileEksporter{"laporan.csv"}}
	//
	//eksportirJson := JsonEksporter{FileEksporter{"laporan.json"}}
	//
	//eksportirKonsol := ConsoleEksporter{}
	//
	//jalankanEksport("laporan keuangan", &eksportirCsv)
	//jalankanEksport("laporan keuangan", &eksportirJson)
	//jalankanEksport("laporan keuangan", &eksportirKonsol)

	//angkaA := 5
	//angkaB := 5
	//
	//fmt.Printf("hasil penjumlahan %d + %d : adalah %d \n", angkaA, angkaB, pertemuan1.Tambah(angkaA, angkaB))
	//fmt.Printf("hasil pengurangan %d + %d : adalah %d \n", angkaA, angkaB, pertemuan1.Kurang(angkaA, angkaB))
	//fmt.Printf("hasil perkalian %d + %d : adalah %d \n", angkaA, angkaB, pertemuan1.Kali(angkaA, angkaB))
	//fmt.Printf("hasil pembagian %d + %d : adalah %d \n", angkaA, angkaB, pertemuan1.Bagi(angkaA, angkaB))
	//
	//hasil, err := pertemuan1.ErrorHandilngfunc(-10)
	//
	//if err != nil {
	//	fmt.Println("Terjadi kesalahan : ", err)
	//} else {
	//	fmt.Println(hasil)
	//}
	//
	//fmt.Println(pertemuan2.ArrayNumber[0])
	//
	//fmt.Println(pertemuan2.SliceNumber)
	//
	//pertemuan2.LearnMap()
	//
	//animal1 := pertemuan2.Animal{Name: "cat", Color: "white", Ovovivivar: true}
	//
	//fmt.Println(animal1)
	//fmt.Println(animal1.Name)
	//fmt.Println(animal1.Color)
	//fmt.Println(animal1.Ovovivivar)
	//
	//pertemuan2.LearnPointer()

	//pertemuan2.AddTodo("Mencuci pakaian")
	//pertemuan2.AddTodo("Makan")
	//pertemuan2.AddTodo("Mandi")
	//pertemuan2.AddTodo("Belajar")

	//for _, todos := range pertemuan2.GetAllTodo() {
	//	fmt.Printf("%v\n", *todos)
	//}

	//getById := pertemuan2.GetTodoById(6)
	//if getById != nil {
	//	fmt.Println(*getById)
	//} else {
	//	fmt.Println("data tidak ditemukan")
	//}

	//err := pertemuan2.Delete(0)
	//if err == nil {
	//	fmt.Println("Data berhasil dihapus")
	//} else {
	//	fmt.Println(err)
	//}
	//
	//for _, todos := range pertemuan2.GetAllTodo() {
	//	fmt.Printf("%v\n", *todos)
	//}

	//tabunganNasabah := pertemuan3.RekeningBank{Nama: "ahmad"}
	//fmt.Println(tabunganNasabah.GetSaldo())
	//
	//tabunganNasabah.Deposit(1000)
	//
	//fmt.Println(tabunganNasabah.GetSaldo())
	//
	//tabunganNasabah.Withdraw(500)
	//
	//fmt.Println(tabunganNasabah.GetSaldo())

}
