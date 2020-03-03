package fdoub

import "math/big"

var (
	zero = new(big.Int)
	one  = big.NewInt(1)
	two  = big.NewInt(2)
)

// A fast Fibonacci calculator through fast doubling.
// https://en.wikipedia.org/wiki/Fibonacci_number#Matrix_form
func Fib(n *big.Int) *big.Int {
	n, _ = fib(n)
	return n
}

func fib(n *big.Int) (*big.Int, *big.Int) {
	if n.Cmp(zero) == 0 {
		return zero, one
	}

	m := new(big.Int)
	n, m = n.DivMod(n, two, m)

	even := m.Cmp(zero) == 0

	a, b := fib(n)

	c, d := m, new(big.Int)
	c = c.Mul(a, a)
	d = d.Mul(b, b)
	d = d.Add(d, c)
	c = c.Mul(b, two)
	c = c.Sub(c, a)
	c = c.Mul(c, a)

	if even {
		return c, d
	}

	c = c.Add(c, d)
	return d, c
}
