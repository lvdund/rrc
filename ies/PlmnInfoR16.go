package ies

import "rrc/utils"

// PLMN-Info-r16 ::= SEQUENCE
type PlmnInfoR16 struct {
	NrBandlistR16 *utils.BITSTRING `lb:maxBandsENDCR16,ub:maxBandsENDCR16`
}
