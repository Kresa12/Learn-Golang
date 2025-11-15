package pengambilandata

import "fmt"

type CreditCardProcessor struct{}

func (c *CreditCardProcessor) PaymentProces(amount float32) error {
	fmt.Printf("Memproses pembayaran kartu kredit sejumlah %.2f\n", amount)
	return nil
}

func (b *CreditCardProcessor) PushNotification(email, message string) error {
	fmt.Printf("Memproses transfer bank sejumlah %.2f\n", email, message)
	return nil
}
