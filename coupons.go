package acm

import (
	"math/big"
)

func Coupons(N int64) (int64, int64, int64) {
	var i int64
	var x, y int64 = 0, 1
	for i = 1; i <= N; i++ {
		x = i*x + y*N
		y = i * y
		gcd := Gcd(x, y)
		x /= gcd
		y /= gcd
		// fmt.Println(i, ":", x, y)
	}

	return x / y, x % y, y
}

func Coupons2(N int64) (int64, int64, int64) {
	var i int64
	n := big.NewRat(0, 1)
	for i = 1; i <= N; i++ {
		n.Add(n, big.NewRat(N, i))
		// fmt.Println(i, ":", n.Num(), n.Denom())
	}
	x, y := n.Num(), n.Denom()
	m := big.NewInt(0)
	x.DivMod(x, y, m)
	return x.Int64(), m.Int64(), y.Int64()
}