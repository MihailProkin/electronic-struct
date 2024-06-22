package main

import (
	"fmt"

	"electronic/electronic"
)

func printCharacteristics(p electronic.Phone) {
	fmt.Printf("Brand: %s\n", p.Brand())
	fmt.Printf("Model: %s\n", p.Model())
	fmt.Printf("Type: %s\n", p.Type())

	if sp, ok := p.(electronic.Smartphone); ok {
		fmt.Printf("OS: %s\n", sp.OS())
	}

	if stp, ok := p.(electronic.StationPhone); ok {
		fmt.Printf("Buttons count: %d\n", stp.ButtonsCount())
	}
}

func main() {
	apple := electronic.NewApplePhone("iPhone 15", "iOS")
	android := electronic.NewAndroidPhone("Samsung", "Galaxy S20", "Android")
	radio := electronic.NewRadioPhone("Panasonic", "KX-UT133", 24)

	printCharacteristics(apple)
	fmt.Println("--------------------")
	printCharacteristics(android)
	fmt.Println("--------------------")
	printCharacteristics(radio)
}
