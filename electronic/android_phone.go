package electronic

// Структура AndroidPhone, реализующая интерфейсы Phone и Smartphone
type AndroidPhone struct {
	Brand string
	Model string
	OS    string
}

func (an AndroidPhone) GetBrand() string {
	return an.Brand
}

func (an AndroidPhone) GetModel() string {
	return an.Model
}

func (an AndroidPhone) Type() string {
	return "Smartphone"
}

func (an AndroidPhone) GetOS() string {
	return an.OS
}
