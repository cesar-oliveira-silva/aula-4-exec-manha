package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	salario := 15001
	err := validaSalario(salario)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("O salario %d necessita pagar imposto \n", salario)
}

func validaSalario(salario int) (erro error) {
	if salario <= 15000 {
		return errors.New("O salario digitado não está dentro do valor mínimo")
	}
	return nil
}
