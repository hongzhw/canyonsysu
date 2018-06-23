package service

import (
	"fmt"
	"testing"
)

func Test_CustomerRegister(t *testing.T) {

	shouldSuccess := []struct {
		name          string
		password      string
		restaurant_id int
		phone         string
		expected      bool
	}{
		{"hzw12", "123", 1, "13719179588", true},
		{"hzw23", "123", 1, "", false},
	}
	for _, ts := range shouldSuccess {
		if real, _ := CustomerRegister(ts.name, ts.password, ts.restaurant_id, ts.phone); real != ts.expected {
			t.Fatalf("fail in " + ts.name + ts.phone)
		} else {
			t.Log("success in " + ts.name + ts.phone)
		}
	}

}

func Test_ListAllCustomers(t *testing.T) {
	real := ListAllCustomers()
	for _, item := range real {
		fmt.Println(item.Name)
	}
	if len(real) == 2 {
		t.Log("success in ")
	} else {
		t.Fatalf("fail in ")
	}
}
