package main

import (
	"fmt"

	"electronic/electronic"
)

func main() {
	// Создаем новый applePhone
	applePhone := electronic.ApplePhone{
		Brand: "Apple",
		Model: "IPhone 15",
		OS:    "iOS",
	}

	fmt.Println("Apple Phone:")
	fmt.Println("Brand:", applePhone.Brand)
	fmt.Println("Model:", applePhone.Model)
	fmt.Println("Type:", applePhone.Type())
	fmt.Println("OS:", applePhone.OS)

	fmt.Println("----------------------------------------------------------------")

	// Создаем новый androidPhone
	androidPhone := electronic.AndroidPhone{
		Brand: "Samsung",
		Model: "Galaxy S20",
		OS:    "Android",
	}

	fmt.Println("Android Phone:")
	fmt.Println("Brand:", androidPhone.Brand)
	fmt.Println("Model:", androidPhone.Model)
	fmt.Println("Type:", androidPhone.Type())
	fmt.Println("OS:", androidPhone.OS)

	fmt.Println("----------------------------------------------------------------")

	// Создаем новый radioPhone
	radioPhone := electronic.RadioPhone{
		Brand:        "Panasonic",
		Model:        "KX-UT133",
		ButtonsCount: 24,
	}

	fmt.Println("Radio Phone:")
	fmt.Println("Brand:", radioPhone.Brand)
	fmt.Println("Model:", radioPhone.Model)
	fmt.Println("Type:", radioPhone.Type())
	fmt.Println("ButtonsCount:", radioPhone.ButtonsCount)

	fmt.Println("----------------------------------------------------------------")
}
