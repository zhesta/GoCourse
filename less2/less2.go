package main

import (
	"fmt"
)

// chet - Проверить четность числа.
func chet(num int) (result bool) {

	if num%2 == 0 {
		result = true
		return result
	}
	return result
}

// three - Проверить деление числа на 3.
func three(num int) (result bool) {

	if num%3 == 0 {
		result = true
		return result
	}
	return result
}

// devision - проверка на простое число.
func devision(delimoe int, delitel int) (result bool) {

	if (delimoe%delitel == 0) && (delimoe != delitel) {
		result = true
		return result
	}
	return result
}

// fiboPrint - Вывод 100 чисел Фибоначи (Fn = Fn-1 + Fn-2).
func fiboPrint() {

	var fibo [100]int
	fibo[2] = 1

	for i := 0; i < 97; i++ {
		fibo[i+3] = fibo[i+2] + fibo[i+1]
	}
	fmt.Println(fibo[2:])
}

// ** Не совсем понял задачу - заполнил массив 2..100 и вычеркнул непростые числа:

// simplePrint - Заполнить массив простыми числами.
func simplePrint() {

	var simple [100]int

	// заполнить массив от 2...100
	for i := 0; i < 100; i++ {
		simple[i] = i + 2
	}

	// Оставить только простые числа
	for p := 2; p < 5; p++ {

		for i := 0; i < len(simple); i++ {

			if devision(simple[i], p) {
				simple[i] = 0
			}
		}
	}
	fmt.Println(simple)
}

func main() {

	c := 3
	fmt.Println("Проверка на деление на 2 = ", chet(c))
	fmt.Println("Проверка на деление на 3 = ", three(c))

	fiboPrint()
	simplePrint()

}
