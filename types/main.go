package main

import "fmt"

func main() {
	var title string
	var copies int
	var name string
	title = "For the Love of Go"
	copies = 99
	name = "Tolkien"
	var edicion string
	edicion = "1 edición"
	var oferta bool
	var descuento int
	oferta = true
	descuento = 15

	fmt.Println(title)
	fmt.Println(copies)
	fmt.Println(name)
	fmt.Println(edicion)
	fmt.Println(oferta)
	fmt.Println(descuento)

	Great("Matias")
	fmt.Println(total([]int{1, 2, 3}))
	fmt.Println("-------")
	evens()
	fmt.Println("-------")

	fmt.Println(noNegative([]int{1, 2, 3, -4, 5, -1, -1, 2, 0, -1}))
	salida()
}

func Great(name string) {
	switch name {
	case "Alice":
		fmt.Printf("Hola, %s", name)
	case "Bob":
		fmt.Printf("Hola, %s", name)
	default:
		fmt.Println("Hola desconocido")
	}

}

func total(numeros []int) int {
	var suma int
	for _, num := range numeros {

		suma = suma + num

	}

	return suma

}

func evens() {
	for x := 0; x < 100; x += 2 {

		fmt.Print(x, ",")
	}
}

func noNegative(numeros []int) []int {
	var lista []int
	for _, n := range numeros {
		if n < 0 {
			continue
		}
		lista = append(lista, n)
	}

	return lista
}
func salida() {
outer:
	for x := range 10 {
		for y := range 10 {
			fmt.Println(x, y)
			if y == 5 {
				continue outer
			}
		}
	}
}
