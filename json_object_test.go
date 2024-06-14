package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName  string
	LastName   string
	Age        int
	Married    bool
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Bayu",
		LastName:   "Sugiantoro",
		Age:        27,
		Married:    false,
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}