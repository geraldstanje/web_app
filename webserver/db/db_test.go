package db

import (
  "testing"
)

func TestUserAdd(t *testing.T) {
  res := AddUser("Max.Musterman@gmail.com", "admin")
  if res != true {
    t.Errorf("AddUser failed")
  }

  res = AddUser("Max.Musterman@gmail.com", "admin")
  if res != false {
    t.Errorf("AddUser failed")
  }
}

func TestCheckUserLogin(t *testing.T) {
  res := AddUser("Kevin_De_Bruyne@gmail.com", "admin")
  if res != true {
    t.Errorf("AddUser failed")
  }

  res = CheckUserLogin("Max.Musterman@gmail.com", "admin")
  if res != true {
    t.Errorf("CheckUserLogin failed")
  }
}