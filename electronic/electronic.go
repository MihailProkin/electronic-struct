package electronic

type Phone interface {
	Brand() string
	Model() string
	Type() string
}

type StationPhone interface {
	ButtonsCount() int
}

type Smartphone interface {
	OS() string
}

// Структура ApplePhone, реализующая интерфейсы Phone и Smartphone
type ApplePhone struct {
	Brand string
	Model string
	OS    string
}

func (a ApplePhone) GetBrand() string {
	return a.Brand
}

func (a ApplePhone) GetModel() string {
	return a.Model
}

func (a ApplePhone) Type() string {
	return "Smartphone"
}

func (a ApplePhone) GetOS() string {
	return a.OS
}

// electronic/electronic.go
// package electronic

// // Phone interface with methods Brand, Model, and Type
// type Phone interface {
//     Brand() string
//     Model() string
//     Type() string
// }

// // StationPhone interface with method ButtonsCount
// type StationPhone interface {
//     ButtonsCount() int
// }

// // Smartphone interface with method OS
// type Smartphone interface {
//     OS() string
// }

// // applePhone struct
// type applePhone struct {
//     brand string
//     model string
//     os    string
// }

// // Implementing Phone interface for applePhone
// func (a applePhone) Brand() string {
//     return a.brand
// }

// func (a applePhone) Model() string {
//     return a.model
// }

// func (a applePhone) Type() string {
//     return "Smartphone"
// }

// // Implementing Smartphone interface for applePhone
// func (a applePhone) OS() string {
//     return a.os
// }

// // Constructor function for applePhone
// func NewApplePhone(brand, model, os string) Phone {
//     return applePhone{
//         brand: brand,
//         model: model,
//         os:    os,
//     }
// }
