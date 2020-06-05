package main

import (
	"fmt"
	"math"
	)

	const (
		 A = 43;
		 B = 25;
	)	

func main () {

	var (
		C float64
		square float64
		perimetr float64
	)	

	square = A * B / 2;
	C = math.Sqrt( math.Pow(A, 2) + math.Pow(B, 2));
	perimetr = A + B + C;

	fmt.Println("Катеты прямоугольника: A = ", A, " B = ", B)
	fmt.Println("Гипотенуза = ", fmt.Sprintf("%.2f",C))
	fmt.Println("Периметр   = ", fmt.Sprintf("%.2f",perimetr))
	fmt.Println("Площадь    = ", square)
}