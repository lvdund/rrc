package ies

import "rrc/utils"

// P0-PUSCH-AlphaSetId ::= utils.INTEGER (0..maxNrofP0-PUSCH-AlphaSets-1)
type P0PuschAlphasetid struct {
	Value utils.INTEGER `lb:0,ub:maxNrofP0PuschAlphasets1`
}
