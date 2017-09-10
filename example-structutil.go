package main

import (
	"fmt"

	"github.com/mariadesouza/utilities/structutil"
)

type testStruct struct {
	Name       string
	Email      []string
	Occupation string
}

func main() {

	testStruct := testStruct{"Ethan", []string{"emdesouza@gmail.com"}, "engineer"}
	m := structutil.StructToMap(&testStruct)
	for key, val := range *m {
		fmt.Println(key, val)
	}
}
