package main

import (
	"errors"
	"fmt"
	"os"
)

type trabalhador struct {
	valorHora int
	horasTrab int
}

func main() {
	//trabalhador 1
	trabalhador1 := trabalhador{valorHora: 1000, horasTrab: 88}
	salario, err := calculaSalarioErrorf(&trabalhador1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("O salario calculado é de %.2f \n", salario)

	//trabalhador 2
	trabalhador2 := trabalhador{valorHora: 1000, horasTrab: 89}
	salario, err = calculaSalarioErrorNew(&trabalhador2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("O salario calculado é de %.2f \n", salario)

	//trabalhador 3
	trabalhador3 := trabalhador{valorHora: 1000, horasTrab: 79}
	salario, err = calculaSalarioErrorUnwrap(&trabalhador3)
	if err != nil {
		fmt.Println(errors.Unwrap(err))
		os.Exit(1)
	}
	fmt.Printf("O salario calculado é de %.2f \n", salario)

}

func calculaSalarioErrorf(colaborador *trabalhador) (salarioc float64, err error) {
	if colaborador.horasTrab < 80 {
		return -1, fmt.Errorf("errorF: o trabalhador não pode ter trabalhado menos de 80 horas por mês. horas trabalhadas: %d", colaborador.horasTrab)
	} else {
		var salario float64
		salario = float64(colaborador.horasTrab * colaborador.valorHora)
		fmt.Printf("salario sem imposto: %.2f\n", salario)
		if salario >= 20000 {
			salario = salario - (salario * 0.1)
			return salario, nil
		}
	}
	return -2, nil

}

func calculaSalarioErrorNew(colaborador *trabalhador) (salarioc float64, err error) {
	if colaborador.horasTrab < 80 {
		return -1, errors.New(fmt.Sprintf("erroNew: o trabalhador não pode ter trabalhado menos de 80 horas por mês. horas trabalhadas: %d", colaborador.horasTrab))
	} else {
		var salario float64
		salario = float64(colaborador.horasTrab * colaborador.valorHora)
		fmt.Printf("salario sem imposto: %.2f\n", salario)
		if salario >= 20000 {
			salario = salario - (salario * 0.1)
			return salario, nil
		}
	}
	return -2, nil

}

func calculaSalarioErrorUnwrap(colaborador *trabalhador) (salarioc float64, err error) {
	//err2 := errorTwo{msg: fmt.Sprintf("erroUnwrap: o trabalhador não pode ter trabalhado menos de 80 horas por mês. horas trabalhadas: %d", colaborador.horasTrab)}
	var err2 = errorTwo{}
	err2.format(fmt.Sprintf("erroUnwrap: o trabalhador não pode ter trabalhado menos de 80 horas por mês. horas trabalhadas: %d", colaborador.horasTrab))
	if colaborador.horasTrab < 80 {
		return -1, fmt.Errorf("Erro: %w \n", err2)
	} else {
		var salario float64
		salario = float64(colaborador.horasTrab * colaborador.valorHora)
		fmt.Printf("salario sem imposto: %.2f\n", salario)
		if salario >= 20000 {
			salario = salario - (salario * 0.1)
			return salario, nil
		}
	}
	return -2, nil

}

type errorTwo struct {
	msg string
}

func (e errorTwo) Error() string {
	return e.msg
}

func (e *errorTwo) format(s string) {
	e.msg = s
}
