package pengambilandata

type PaymentProcessor interface {
	PaymentProces(amount float32) error
}
