package db

import (
  "testing"
)

func TestUserAdd(t *testing.T) {
  res := AddUser("Max.Musterman@gmail.com", "admin")
  if res != true {
    t.Errorf("Add user failed")
  }

  res = AddUser("Max.Musterman@gmail.com", "admin")
  if res != false {
    t.Errorf("Add user failed")
  }
}