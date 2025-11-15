package pengambilandata

import "fmt"

type OrderService struct {
	payment      PaymentProcessor
	notification NotificationService
}

func NewOrderService(payment PaymentProcessor, notification NotificationService) *OrderService {
	return &OrderService{
		payment,
		notification,
	}
}

func (o *OrderService) OrderProces(order *Order) {
	err := o.payment.PaymentProces(order.Amount)
	if err != nil {
		fmt.Printf("Pesanan Anda Gagal : %v \n", err)
	}
	message := fmt.Sprintf("Pesanan Anda #%s telah berhasil.", order.Id)
	err = o.notification.PushNotification(order.EmailUser, message)
	if err != nil {
		fmt.Printf("Pembayaran berhasil, namun notifikasi gagal : %v \n", err)
	} else {
		fmt.Printf("Pesanan berhasil diproses!\n")
	}
}
