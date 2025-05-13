package main

import (
	"creditcalculator/internal/calculator"
	"creditcalculator/internal/cli"
	"fmt"
)

func main() {

	var menuNumbers int8
	var flag bool = false

	fmt.Printf("Добро пожаловать в кредитный калькулятор.\n")

	for flag != true {

		menuNumbers = int8(cli.ReadInt("\n_________________________________________\n\nВыберите действие: \n 1. Расчитать платеж по аннуитетному кредиту \n10. Выход из программы\n______________________________________\n"))
		switch menuNumbers {
		case 1:
			{
				fmt.Printf("Введите сумму кредита => ")
				sumCredit := cli.ReadFloat("Введите сумму кредита => ")
				term := cli.ReadInt("Введите срок кредита (кол-во месяцев) => ")
				rate := cli.ReadFloat("Введите процентную ставку => ")
				monthlyPayment := calculator.CalculateAnnuityPayment(int(term), rate, sumCredit)
				fmt.Printf("\n_________________________________________\nЕжемесячный платеж по кредиту будет составлять => %.2f", monthlyPayment)
			}
		case 10:
			{
				flag = true
			}
		}
	}

}
