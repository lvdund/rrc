package ies

import "rrc/utils"

// BandSidelinkPC5-r16-sl-Reception-r16-scs-CP-PatternRxSidelink-r16-fr1-r16 ::= SEQUENCE
type Bandsidelinkpc5R16SlReceptionR16ScsCpPatternrxsidelinkR16Fr1R16 struct {
	Scs15khzR16 *utils.BITSTRING `lb:16,ub:16`
	Scs30khzR16 *utils.BITSTRING `lb:16,ub:16`
	Scs60khzR16 *utils.BITSTRING `lb:16,ub:16`
}
