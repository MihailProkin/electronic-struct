package main

import (
	"fmt"

	"github.com/mihailprokin/electronic-struct/electronic"
)

func main() {
	applePhone := electronic.ApplePhone{
		Brand: "Apple",
		Model: "IPhone 15",
		OS:    "iOS",
	}

	fmt.Println("Apple Phone:")
	fmt.Println("Brand: ", applePhone.Brand())
	fmt.Println("Model: ", applePhone.Model())
	fmt.Println("Type: ", applePhone.Type())
	fmt.Println("OS: ", applePhone.OS())
}

// main.go
// package main

// import (
//     "fmt"
//     "your_project_path/electronic"
// )

// func main() {
//     // Создаем новый applePhone
//     myPhone := electronic.NewApplePhone("Apple", "iPhone 13", "iOS")

//     // Используем методы интерфейсов
//     fmt.Println("Brand:", myPhone.Brand())
//     fmt.Println("Model:", myPhone.Model())
//     fmt.Println("Type:", myPhone.Type())

//     // Приводим к интерфейсу Smartphone, чтобы использовать метод OS
//     if smartphone, ok := myPhone.(electronic.Smartphone); ok {
//         fmt.Println("OS:", smartphone.OS())
//     } else {
//         fmt.Println("This phone is not a smartphone")
//     }
// }
