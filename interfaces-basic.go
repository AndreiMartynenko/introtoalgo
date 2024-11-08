package main

/*

// define an interface
type Payer interface {
	Pay(int) error
}

// struct Wallet
type Wallet struct {
	Cash int
}

// Method for Wallet
func (w *Wallet) Pay(amount int) error {
	if w.Cash < amount {
		return fmt.Errorf("Not enough money")
	}
	w.Cash -= amount
	return nil
}

// function Buy that takes a Payer interface
func Buy(p Payer) {
	err := p.Pay(10)
	if err != nil {
		panic(err)
	}
	//The %T verb prints the type of the value passed to it.
	fmt.Printf("Thank you for your purchase %T\n\n", p)
}

func main() {
	myWallet := &Wallet{Cash: 100}
	Buy(myWallet)
}

*/
