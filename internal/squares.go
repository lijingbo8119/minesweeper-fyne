package internal

type Squares []*Square

func (this Squares) find(closure func(s *Square) bool) *Square {
	for _, s := range this {
		if closure(s) {
			return s
		}
	}
	return nil
}

func (this Squares) filter(closure func(s *Square) bool) Squares {
	res := Squares{}
	for _, s := range this {
		if closure(s) {
			res = append(res, s)
		}
	}
	return res
}

func (this Squares) each(closure func(s *Square)) {
	for _, s := range this {
		closure(s)
	}
}
