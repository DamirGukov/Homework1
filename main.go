package main

import "fmt"

func main() {
	var number1 float64 = 54.545
	number2 := false
	name := "John"

	fmt.Printf(" %.2f\n ", number1)                  // для виведення чисел із плаваючою точкою
	fmt.Printf("%T\n", number2)                      // для виведення типу змінної
	fmt.Printf("%E\n", number1)                      // для виведення чисел з плаваючою точкою в експоненційному вигляді
	fmt.Printf("%s", name)                           // для виведення рядка
	fmt.Printf("%v %.2v %v", name, number1, number2) // універсальний специфікатор %v
}
