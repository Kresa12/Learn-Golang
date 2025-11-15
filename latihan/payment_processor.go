package latihan

type PaymentProcessor interface {
	PaymentProces(amount float32) error
}
