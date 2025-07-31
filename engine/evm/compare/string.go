package compare

type String struct {
	Falsy[string]
	inner string
}

func NewString(raw string) (Compare[string], error) {
	return &String{
		inner: raw,
	}, nil
}

func (s *String) Fixed(v string) bool {
	return s.inner == v
}
