package dpd

type Dpd struct {
	URL    string
	CLIENT string
	KEY    string
}

func (d *Dpd) GetPoints(params []byte) ([]byte, error) {

	return []byte(""), nil
}

func (d *Dpd) GetCourier(params []byte) ([]byte, error) {

	return []byte(""), nil
}
