package latihan

import "fmt"

type CreditCardProcessor struct{}

func (c *CreditCardProcessor) PaymentProces(amount float32) error {
	fmt.Printf("Memproses pembayaran kartu kredit sejumlah %.2f\n", amount)
	return nil
}
