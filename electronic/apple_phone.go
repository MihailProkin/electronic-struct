package electronic

// Структура ApplePhone, реализующая интерфейсы Phone и Smartphone
type ApplePhone struct {
	Brand string
	Model string
	OS    string
}

func (ap ApplePhone) GetBrand() string {
	return ap.Brand
}

func (ap ApplePhone) GetModel() string {
	return ap.Model
}

func (ap ApplePhone) Type() string {
	return "Smartphone"
}

func (ap ApplePhone) GetOS() string {
	return ap.OS
}
