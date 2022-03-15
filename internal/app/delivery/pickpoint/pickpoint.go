package pickpoint

type Pickpoint struct {
	URL   string
	LOGIN string
	PASS  string
	IKN   string
}

func (p *Pickpoint) GetPoints(params []byte) ([]byte, error) {

	return []byte(""), nil
}

func (p *Pickpoint) GetCourier(params []byte) ([]byte, error) {

	return []byte(""), nil
}
