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
