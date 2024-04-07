package main

import (
	"fmt"
	"os"
)

type erroSalario struct {
	status  int
	msgErro string
}

type error interface {
	Error() string
}

func main() {

	//var salario int

	salario := 14000
	status, err := validaSalario(salario)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Status %d. O salario %d necessita pagar imposto \n", status, salario)
}
func (e *erroSalario) Error() string {
	return fmt.Sprintf("%d - %v", e.status, e.msgErro)
}

func validaSalario(salario int) (int, error) {
	if salario <= 15000 {
		return 500, &erroSalario{
			status:  500,
			msgErro: "O salario digitado não está dentro do valor mínimo",
		}
	}
	return 200, nil
}
