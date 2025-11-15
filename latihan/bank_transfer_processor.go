package latihan

import "fmt"

type BankTransferProcessor struct{}

func (b *BankTransferProcessor) PaymentProces(amount float32) error {
	fmt.Printf("Memproses transfer bank sejumlah %.2f\n", amount)
	return nil
}
