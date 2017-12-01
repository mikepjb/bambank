package main

import (
  "fmt"
  "time"
)

type Account struct {
  id int
  name string
  password_encrypted string
}

type Transaction struct {
  id int
  timestamp time.Time
  credit_id int
  debit_id int
  amount int
}

var Transactions []Transaction
var Accounts []Account

func createAccount(id int, name string, password string) Account {
  account := Account{id, name, password} // TODO way to increment id?
  Accounts = append(Accounts, account)
  Transactions = append(Transactions, Transaction{1, time.Now(), 1, account.id, 100})
  return account
}

func (account Account) debit(receiving Account, amount int) {
  Transactions = append(Transactions, Transaction{1, time.Now(), receiving.id, account.id, amount})
}

func (account Account) balance() int {
  balance := 0
  for _, transaction := range Transactions {
    if transaction.credit_id == account.id {
      balance -= transaction.amount
    } else if transaction.debit_id == account.id {
      balance += transaction.amount
    }
  }
  return balance
}

func main() {
  masterAccount := Account{1, "Master Account", "letmein"}
  Accounts = append(Accounts, masterAccount)

  testAccountOne := createAccount(5, "holy", "diver")
  testAccountTwo := createAccount(7, "unholy", "climber")
  testAccountOne.debit(testAccountTwo, 50)

  fmt.Printf("one: %v\n", testAccountOne.balance())
  fmt.Printf("two: %v\n", testAccountTwo.balance())
  fmt.Printf("master: %v\n", masterAccount.balance())

  fmt.Println(Transactions)
  fmt.Println(Accounts)
}
