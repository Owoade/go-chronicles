package main

import "fmt"

type PaymentMethod interface {
	Pay() string
}

type CardPaymentMethod struct {
	number int
	exp string
	cvv string
}

func (m CardPaymentMethod) Pay() string{
	return fmt.Sprintf("payment made using card details %d, %s, %s", m.number, m.exp, m.cvv)
}

type BankTransferPaymentMethod struct {
	accountNumber int
	bankName string
}

func (m BankTransferPaymentMethod) Pay() string {
	return fmt.Sprintf("payment made using bank transfer from %d, %s", m.accountNumber, m.bankName)
}

func MakePayment(m PaymentMethod){
	fmt.Println(m.Pay())
}

func main(){
	card := CardPaymentMethod{
		number: 12122120090992949,
		exp: "02/28",
		cvv: "009",
	}

	transfer := BankTransferPaymentMethod{
		accountNumber: 2121117368,
		bankName: "UNITED BANK OF AFRICA",
	}

	MakePayment(card)
	MakePayment(transfer)
}

