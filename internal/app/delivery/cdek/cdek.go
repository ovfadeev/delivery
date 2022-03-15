package cdek

type Cdek struct {
	URL   string
	LOGIN string
}

func (c *Cdek) GetPoints(params []byte) ([]byte, error) {

	return []byte(""), nil
}

func (c *Cdek) GetCourier(params []byte) ([]byte, error) {

	return []byte(""), nil
}
