package electronic

// Структура ApplePhone, реализующая интерфейсы Phone и Smartphone
type applePhone struct {
	model string
	os    string
}

func NewApplePhone(model, os string) *applePhone {
	return &applePhone{
		model: model,
		os:    os,
	}
}

func (ap *applePhone) Brand() string {
	return "Apple"
}

func (ap *applePhone) Model() string {
	return ap.model
}

func (ap *applePhone) Type() string {
	return "smartphone"
}

func (ap *applePhone) OS() string {
	return ap.os
}
