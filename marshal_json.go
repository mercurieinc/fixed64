// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 16:42:36 2644A5              priveda/fixed64/[marshal_json.go]
// -----------------------------------------------------------------------------

package fixed64

import (
	"fmt"
	"strings"
)

// MarshalJSON returns the JSON encoding of zr.Fixed64.
func (ob Fixed64) MarshalJSON() ([]byte, error) {
	//
	// TODO: using fmt.Sprintf() may slow down performance.
	//       There are faster ways to build a number with 4 decimals.
	//       Create a benchmark to find the fastest method.
	//
	var (
		i   = ob.i64 / 1E4   // integer part
		d   = ob.i64 - i*1E4 // decimal part
		ret = fmt.Sprintf("%d", i)
	)
	if d != 0 {
		ret += strings.TrimRight(
			fmt.Sprintf("%0.4f", float32(d)/1E4)[1:],
			"0",
		)
	}
	return []byte(ret), nil
}

//end
