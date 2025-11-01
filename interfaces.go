package main 
import (
	"fmt"
)

type paymenter interface {
	pay(amount float32)        // executes the method behind every type
}

type payment struct {
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	// razorpayPaymentGw := razorpay{}
	// razorpayPaymentGw.pay(amount)
	p.gateway.pay(amount)
	

}

type razorpay struct {

}

func (r razorpay) pay(amount float32) {
	//logic to pay payment

	fmt.Println("paid using razorpay", amount)
}

type stripe struct {

}

func (s stripe) pay(amount float32) {
	fmt.Println("paid using stripe: ", amount)
}

type paypal struct {

}

func (p paypal) pay(amount float32) {
	fmt.Println("made payment using paypal:", amount)
}


func main(){
	fmt.Println("")
	paypalPaymentGw := paypal{}
	newPayment := payment{
		gateway: paypalPaymentGw,        // interface sees the method behind a struct, executes it and all happy!
	}

	newPayment.makePayment(100.59)


}
