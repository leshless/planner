package xslices

func Map[A any, B any](as []A, mapping func(a A) B) []B {
	bs := make([]B, 0, len(as))
	for _, a := range as {
		bs = append(bs, mapping(a))
	}

	return bs
}
