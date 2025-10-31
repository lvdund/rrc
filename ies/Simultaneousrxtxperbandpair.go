package ies

import "rrc/utils"

// SimultaneousRxTxPerBandPair ::= BIT STRING (SIZE (3..496))
type Simultaneousrxtxperbandpair struct {
	Value utils.BITSTRING `lb:3,ub:496`
}
