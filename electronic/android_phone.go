package electronic

// Структура AndroidPhone, реализующая интерфейсы Phone и Smartphone
type androidPhone struct {
	brand string
	model string
	os    string
}

func NewAndroidPhone(brand, model, os string) *androidPhone {
	return &androidPhone{
		brand: brand,
		model: model,
		os:    os,
	}
}

func (an *androidPhone) Brand() string {
	return an.brand
}

func (an *androidPhone) Model() string {
	return an.model
}

func (an *androidPhone) Type() string {
	return "smartphone"
}

func (an *androidPhone) OS() string {
	return an.os
}
