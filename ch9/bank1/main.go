package main

import (
	"fmt"
)

type withdrawCh struct {
	amount    int
	isSuccess chan bool
}

var deposits = make(chan int)
var balances = make(chan int)
var withdrawals = make(chan *withdrawCh)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	isSuccess := make(chan bool)
	withdrawals <- &withdrawCh{amount: amount, isSuccess: isSuccess}
	return <-isSuccess
}

// The teller is the monitor goroutine
func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case withdrawal := <-withdrawals:
			tempBalance := balance - withdrawal.amount
			if tempBalance > 0 {
				balance = tempBalance
				withdrawal.isSuccess <- true
			} else {
				withdrawal.isSuccess <- false
			}

		case balances <- balance:
		}
	}
}

func init() {
	go teller()
}

func main() {
	Deposit(200)
	fmt.Println(Balance())
	Deposit(2000)
	fmt.Println(Balance())
	Deposit(20)
	fmt.Println(Balance())
	fmt.Println(Balance())
	fmt.Println(Balance())
	fmt.Println(Balance())
	fmt.Println(Balance())
	fmt.Println(Withdraw(40))
	fmt.Println(Balance())
	fmt.Println(Withdraw(400))
	fmt.Println(Balance())
	fmt.Println(Withdraw(4000))
	fmt.Println(Balance())

}
