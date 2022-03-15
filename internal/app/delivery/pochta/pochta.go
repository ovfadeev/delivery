package pochta

type Pochta struct {
	URL     string
	ACCOUNT string
}

func (p *Pochta) GetPoints(params []byte) ([]byte, error) {

	return []byte(""), nil
}

func (p *Pochta) GetCourier(params []byte) ([]byte, error) {

	return []byte(""), nil
}
