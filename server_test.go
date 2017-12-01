package main

import (
  "testing"
)

func TestAccountCreationInsertsTransaction(t *testing.T) {
  if len(Transactions) != 0 {
    t.Error("expected 0")
  }

  createAccount(47, "testy", "tester")

  if len(Transactions) != 1 {
    t.Error("expected 1")
  }
}

func TestTransactionsAffectBothParties(t *testing.T) {
  testAccountOne := createAccount(5, "holy", "diver")
  testAccountTwo := createAccount(7, "unholy", "climber")

  if testAccountOne.balance() != 100 {
    t.Error("expected testAccountOne's balance to be 100")
  }

  if testAccountTwo.balance() != 100 {
    t.Error("expected testAccountTwo's balance to be 100")
  }

  testAccountOne.debit(testAccountTwo, 50)

  if testAccountOne.balance() != 50 {
    t.Error("expected testAccountOne's balance to be 50")
  }

  if testAccountTwo.balance() != 150 {
    t.Error("expected testAccountTwo's balance to be 150")
  }
}
