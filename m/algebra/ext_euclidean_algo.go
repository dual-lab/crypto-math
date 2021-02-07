package algebra

func ExtEuclideanAlgo(a, b int, unsinged bool) (g, u, v int) {
	if a < b {
		a, b = b, a // switch if a < b
	}
	u = 1
	g = a
	if b == 0 {
		return g, u, 0 // handle b=0 case.
	}
	x := 0
	r := b
	for r > 0 {
		// euclidean algo step
		q := g / r
		t := g - q*r
		// u coefficent from euclidean algo
		s := u - q*x
		u, g, x, r = x, r, s, t
	}
	if unsinged && u < 0 {
		// Force u positive
		u = u + (-u*g/b+1)*(b/g)
	}

	v = (g - a*u) / b
	return
}
