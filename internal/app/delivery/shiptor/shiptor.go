package shiptor

type Shiptor struct {
	URL string
	KEY string
}

func (s *Shiptor) GetPoints(params []byte) ([]byte, error) {

	return []byte(""), nil
}

func (s *Shiptor) GetCourier(params []byte) ([]byte, error) {

	return []byte(""), nil
}
