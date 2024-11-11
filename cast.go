package main

import (
	"fmt"
)

// --------------

type Wallet struct {
	Cash int
}

func (w *Wallet) Pay(amount int) error {
	if w.Cash < amount {
		return fmt.Errorf("Not enough cash")
	}
	w.Cash -= amount
	return nil
}

// --------------

type Card struct {
	Balance    int
	ValidUntil string
	Cardholder string
	CVV        string
	Number     string
}

func (c *Card) Pay(amount int) error {
	if c.Balance < amount {
		return fmt.Errorf("Not enough money on balance")
	}
	c.Balance -= amount
	return nil
}

// --------------

type ApplePay struct {
	Money   int
	AppleID string
}

func (a *ApplePay) Pay(amount int) error {
	if a.Money < amount {
		return fmt.Errorf("Not enough money on account")
	}
	a.Money -= amount
	return nil
}

// --------------

type Payer interface {
	Pay(int) error
}

// --------------

func Buy(p Payer) {
	switch p.(type) {
	case *Wallet:
		fmt.Println("Cash payment?")
	case *Card:
		plasticCard, ok := p.(*Card)
		if !ok {
			fmt.Println("Failed to cast to Card")
		}
		fmt.Println("Card payment? Cardholder:", plasticCard.Cardholder)
	default:
		fmt.Println("Unknown payment method")
	}

	err := p.Pay(10)
	if err != nil {
		fmt.Printf("Failed to perform payment %T: %v\n", p, err)
		return
	}
	fmt.Printf("Thank you for your purchase %T\n\n", p)
}

// --------------

// func main() {

// 	myWallet := &Wallet{Cash: 100}
// 	Buy(myWallet)

// 	var myMoney Payer
// 	myMoney = &Card{Balance: 100, Cardholder: "Alice T."}
// 	Buy(myMoney)

// 	myMoney = &ApplePay{Money: 9}
// 	Buy(myMoney)
// }
