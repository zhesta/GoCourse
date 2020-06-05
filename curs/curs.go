package main

import (
	"fmt"
	)

	const cursUSD = 74.32;

func main () {
	var rub float64
	fmt.Println("Введите кол-во рублей:")
	fmt.Scanln(&rub)
	fmt.Println("Введено:",rub,"руб. Эквивалент:", fmt.Sprintf("%.2f",rub / cursUSD),"USD")

}