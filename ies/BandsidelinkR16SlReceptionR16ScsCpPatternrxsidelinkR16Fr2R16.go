package ies

import "rrc/utils"

// BandSidelink-r16-sl-Reception-r16-scs-CP-PatternRxSidelink-r16-fr2-r16 ::= SEQUENCE
type BandsidelinkR16SlReceptionR16ScsCpPatternrxsidelinkR16Fr2R16 struct {
	Scs60khzR16  *utils.BITSTRING `lb:16,ub:16`
	Scs120khzR16 *utils.BITSTRING `lb:16,ub:16`
}
