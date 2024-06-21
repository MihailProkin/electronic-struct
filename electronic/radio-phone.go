package electronic

// Структура RadioPhone, реализующая интерфейсы Phone и StationPhone
type RadioPhone struct {
	Brand        string
	Model        string
	ButtonsCount int
}

func (r RadioPhone) GetBrand() string {
	return r.Brand
}

func (r RadioPhone) GetModel() string {
	return r.Model
}

func (r RadioPhone) Type() string {
	return "Station"
}

func (r RadioPhone) GetButtonsCount() int {
	return r.ButtonsCount
}
