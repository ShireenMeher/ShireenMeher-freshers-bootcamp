package main

import (
	"fmt"
	"sync"
)

type bank_account struct {
	balance int
	mtx sync.Mutex
}

func withdraw(bank *bank_account, money int, wg *sync.WaitGroup){
	defer wg.Done()
	bank.mtx.Lock()
		if(bank.balance-money<0) {
			fmt.Println("withdrawal not possible")
			bank.mtx.Unlock()
			return
		}


		bank.balance=bank.balance-money
		fmt.Println("the new bank balance is", bank.balance)
		bank.mtx.Unlock()
}

func deposit(bank *bank_account, money int, wg *sync.WaitGroup){
	defer wg.Done()
	bank.mtx.Lock()
	bank.balance=bank.balance+money
	fmt.Println("the new bank balance is", bank.balance)
	bank.mtx.Unlock()
}

func main(){
	wg:=new (sync.WaitGroup)
	wg.Add(3)
	bank := bank_account{500, *new(sync.Mutex)}

	go deposit(&bank,500, wg )
	go withdraw(&bank, 700 ,wg)
	go withdraw(&bank, 500 ,wg )

	wg.Wait()
}