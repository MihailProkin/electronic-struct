package electronic

// Структура RadioPhone, реализующая интерфейсы Phone и StationPhone
type radioPhone struct {
	brand        string
	model        string
	buttonsCount int
}

func NewRadioPhone(brand, model string, buttonsCount int) *radioPhone {
	return &radioPhone{
		brand:        brand,
		model:        model,
		buttonsCount: buttonsCount,
	}
}

func (r *radioPhone) Brand() string {
	return r.brand
}

func (r *radioPhone) Model() string {
	return r.model
}

func (r *radioPhone) Type() string {
	return "station"
}

func (r *radioPhone) ButtonsCount() int {
	return r.buttonsCount
}
