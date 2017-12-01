package main

import (
  "fmt"
  "time"
  "net/http"
  "os"
  "log"
  "html/template"
  "strings"
  "strconv"
)

type Account struct {
  id int
  name string
  password string
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
  account := Account{id, name, password}
  Accounts = append(Accounts, account)
  Transactions = append(Transactions, Transaction{1, time.Now(), 1, account.id, 100})
  return account
}

func (account Account) debit(receiving Account, amount int) {
  Transactions = append(Transactions, Transaction{1, time.Now(), account.id, receiving.id, amount})
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

func prepareIndexPage(w http.ResponseWriter, data map[string]string) {
  filepath := "site/index.html"
  fmt.Println(filepath)
  fmt.Println(Transactions)
  t, _ := template.ParseFiles(filepath)
  t.Execute(w, data)
}

func main() {
  masterAccount := Account{1, "Master Account", "letmein"}
  Accounts = append(Accounts, masterAccount)

  createAccount(47, "testuser", "testcity")
  createAccount(77, "anothertestuser", "testcity")

  data := map[string]string{}

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    prepareIndexPage(w, data)
  })

  http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
      r.ParseForm()
      for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
      }

      for _, account := range Accounts {
        if r.FormValue("username") == account.name {
          fmt.Printf("found match for username %s\n", account.name)
          if r.FormValue("password") == account.password {
            fmt.Printf("found match for password %s\n", account.password)
            data["username"] = account.name
            data["loggedIn"] = "yes"
            data["balance"] = strconv.Itoa(account.balance())
          }
        }
      }

      prepareIndexPage(w, data)
    }
  })

  http.HandleFunc("/debit", func(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
      r.ParseForm()

      var credit_account Account
      var debit_account Account

      for _, account := range Accounts {
        if r.FormValue("creditor") == account.name {
          credit_account = account
        } else if r.FormValue("username") == account.name {
          debit_account = account
        }
      }

      amount, _ := strconv.Atoi(r.FormValue("amount"))
      credit_account.debit(debit_account, amount)

      data["balance"] = strconv.Itoa(credit_account.balance())
      prepareIndexPage(w, data)
    }
  })

  http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
      data = map[string]string{}
      prepareIndexPage(w, data)
    }
  })

  port := os.Getenv("PORT")

  if port == "" {
    port = "8080"
  }

  fmt.Printf("Starting serving on port %s\n", port)
  log.Fatal(http.ListenAndServe(":"+port, nil))
}
