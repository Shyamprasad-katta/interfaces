package main

import (
	"fmt"
)

type PaymentGateway interface {
	Charge(amount float64) error
	Refund(transactionID string) error
}

type StripeGateway struct {
	APIKey string
}

func (s StripeGateway) Charge(amount float64) error {
	fmt.Printf("Stripe: Charging $%.2f\n", amount)
	return nil
}

func (s StripeGateway) Refund(tx string) error {
	fmt.Printf("Stripe: Refunding transaction %s\n", tx)
	return nil
}

type PaypalGateway struct {
	ClientID string
}

func (p PaypalGateway) Charge(amount float64) error {
	fmt.Printf("PayPal: Charging $%.2f\n", amount)
	return nil
}

func (p PaypalGateway) Refund(tx string) error {
	fmt.Printf("PayPal: Refunding transaction %s\n", tx)
	return nil
}

func Checkout(gateway PaymentGateway, amount float64) {

	err := gateway.Charge(amount)
	if err != nil {
		fmt.Println("Payment Failed")
		return
	}

	fmt.Println("Payment Successful")
}

func Refund(gateway PaymentGateway, tx string) {
	err := gateway.Refund(tx)
	if err != nil {
		fmt.Println("Refund Failed")
		return
	}

	fmt.Println("Refund Successful")
}

func main() {
	stripe := StripeGateway{
		APIKey: "stripe-secret-key",
	}
	paypal := PaypalGateway{
		ClientID: "paypal-client",
	}
	Checkout(stripe, 100)
	Checkout(paypal, 250)
	Refund(stripe, "TX1001")
	Refund(paypal, "TX2001")
}
