package redexpress

type Redexpress struct {
	URL   string
	LOGIN string
	PASS  string
}

func (r *Redexpress) GetPoints(params []byte) ([]byte, error) {

	return []byte(""), nil
}

func (r *Redexpress) GetCourier(params []byte) ([]byte, error) {

	return []byte(""), nil
}
