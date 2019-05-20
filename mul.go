// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-20 01:32:42 D311E8                       priveda/fixed64/[mul.go]
// -----------------------------------------------------------------------------

package fixed64

import (
	"math/big"
)

// Mul multiplies a fixed-point number by one or more fixed-point
// numbers and returns the result. The original number is not changed.
func (n Fixed64) Mul(nums ...Fixed64) Fixed64 {
	for _, num := range nums {
		var (
			a = n.i64
			b = num.i64
		)
		// return zero if either number is zero
		if a == 0 || b == 0 {
			return Fixed64{0}
		}
		// if direct multiplication will overflow int64, use big.Int
		limit := maxInt64 / a
		if limit < 0 {
			limit = -limit
		}
		if b >= limit || b <= -limit {
			c := big.NewInt(a)
			c.Mul(c, big.NewInt(b))
			c.Div(c, big1E4)
			//
			// if result can't be stored in Fixed64, return overflow
			//
			// TODO: IsInt64() is not available in older Go versions ``
			overflow := !c.IsInt64()
			var ret int64
			if !overflow {
				ret = c.Int64()
				if ret < minInt64 || ret > maxInt64 {
					overflow = true
				}
			}
			if overflow {
				return fixed64Overflow(
					(a < 0 || b < 0) && (a > 0 || b > 0),
					a, " * ", b, " = ", c)
			}
			return Fixed64{ret}
		}
		n.i64 = a * b / 1E4
	}
	return n
}

//end
