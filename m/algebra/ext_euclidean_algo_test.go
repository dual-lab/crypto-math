package algebra_test

import (
	"fmt"
	"testing"

	"com.github/dual-lab/crypto-math/m/algebra"
)

////////////////////////////////////////////
// GCD test + bench
////////////////////////////////////////////
var tableTestGCD = [...]struct{ a, b, g, u, v int }{
	{a: 527, b: 1258, g: 17, u: 13, v: -31},
	{a: 228, b: 1056, g: 12, u: 8, v: -37},
	{a: 163961, b: 167181, g: 7, u: -4430, v: 4517},
	{a: 3892394, b: 239847, g: 1, u: 59789, v: -970295},
	{a: 3, b: 0, g: 3, u: 1, v: 0},
}

func TestExtGCD(t *testing.T) {
	for idx, tCase := range tableTestGCD {
		var tName = fmt.Sprintf("%d. au+bv=gcd(%d,%d)", idx, tCase.a, tCase.b)
		t.Run(tName, func(t *testing.T) {
			g, u, v := algebra.ExtEuclideanAlgo(tCase.a, tCase.b, false)
			if g != tCase.g {
				t.Errorf("GCD expeced %d -- actual %d", tCase.g, g)
			}

			if u != tCase.u {
				t.Errorf("u coeff. expeced %d -- actual %d", tCase.u, u)
			}

			if g != tCase.g {
				t.Errorf("v coeff. expeced %d -- actual %d", tCase.v, v)
			}
		})
	}
}

func TestExtGCDUPositive(t *testing.T) {
	for idx, tCase := range tableTestGCD {
		var tName = fmt.Sprintf("%d. au+bv=gcd(%d,%d)", idx, tCase.a, tCase.b)
		t.Run(tName, func(t *testing.T) {
			_, u, _ := algebra.ExtEuclideanAlgo(tCase.a, tCase.b, true)
			if u < 0 {
				t.Errorf("u coefficent expeced to be positive -- actual %d", u)
			}
		})
	}
}

func BenchmarkExtGCD(b *testing.B) {
	for idx, tCase := range tableTestGCD {
		var tName = fmt.Sprintf("%d. au+bv=gcd(%d,%d)", idx, tCase.a, tCase.b)
		b.Run(tName, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				algebra.ExtEuclideanAlgo(tCase.a, tCase.b, false)
			}
		})
	}
}

func BenchmarkExtGCDUPositive(b *testing.B) {
	for idx, tCase := range tableTestGCD {
		var tName = fmt.Sprintf("%d. au+bv=gcd(%d,%d)", idx, tCase.a, tCase.b)
		b.Run(tName, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				algebra.ExtEuclideanAlgo(tCase.a, tCase.b, true)
			}
		})
	}
}

////////////////////////////////////////////
// Inverse mod m test + bench
////////////////////////////////////////////
var tableTestInvModM = [...]struct{ a, m, r int }{
	{a: 2, m: 5, r: 3},
	{a: 25, m: 7, r: 2},
	{a: 3, m: 8, r: 3},
}

func TestExtInvModM(t *testing.T) {
	for idx, tCase := range tableTestInvModM {
		var tName = fmt.Sprintf("%d. a%d=1 mod %d", idx, tCase.a, tCase.m)
		t.Run(tName, func(t *testing.T) {
			inv, _ := algebra.InverModM(tCase.a, tCase.m)
			if inv != tCase.r {
				t.Errorf("Inverse expeced %d -- actual %d", tCase.r, inv)
			}
		})
	}
}
