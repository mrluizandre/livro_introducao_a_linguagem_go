package main

import "fmt"

func main() {
	// fmt.Println("1 + 1 = ", 1.0+1.0)

	// fmt.Println(len("Andre"))
	// fmt.Println("Andre"[1])
	// fmt.Println("Andre " + "1")

	// var name string = "Andre"
	// var last_name string
	// middle_name := "Luiz"
	// var other_middle_name = "Cardoso da"
	// age := 27
	// var height = 1.76

	// last_name = "Costa"

	// fmt.Println(name, middle_name, other_middle_name, last_name)
	// fmt.Println(age)
	// fmt.Println(height)

	// var (
	// 	a = 1
	// 	b = 2
	// 	c = "batata"
	// )

	// fmt.Println(a, b, c)

	// // fmt.Println(ftoc())

	// for i := 0; i <= 10; i++ {
	// 	switch i {
	// 	case 0:
	// 		fmt.Println("Zero")
	// 	case 1:
	// 		fmt.Println("Um")
	// 		fmt.Println("One")
	// 	default:
	// 		fmt.Println("Mais que um")
	// 	}
	// }
	i := 0
	for i < 1000000000 {
		i++
	}
	fmt.Println(i)
}

func ftoc() float64 {
	var tempF float64

	fmt.Printf("Informe a temperatura em Fahrenheit: ")
	fmt.Scanf("%f", &tempF)

	return (tempF - 32) * 5 / 9
}
