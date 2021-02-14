package algebra

import (
	"fmt"
)

func ExtEuclideanAlgo(a, b int, unsinged bool) (g, u, v int) {
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

func ExampleExtEuclideanAlgo() {
	g, _, _ := ExtEuclideanAlgo(55, 25, false)
	fmt.Printf("gcd(55, 25)=%d", g)
	// Output:
	// 5
}

func InverModM(a, m int) (int, bool) {
	var argsSwitched bool
	if a < m {
		a, m = m, a
		argsSwitched = true
	}

	g, u, v := ExtEuclideanAlgo(a, m, true)
	// inv a mod m existis if and only if gcd(a,m) = 1
	if g != 1 {
		return 0, false
	}
	var inv int = u
	if argsSwitched {
		if v < 0 {
			v = v - (v*g/a-1)*(a/g)
		}
		inv = v
	}

	return inv, true
}

func ExampleInverModM() {
	inv, _ := InverModM(2, 5)
	fmt.Printf("%d*%d=1 mod %d", 2, inv, 5)
}
