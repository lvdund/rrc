package ies

import "rrc/utils"

// SRS-Resource-freqHopping ::= SEQUENCE
type SrsResourceFreqhopping struct {
	CSrs utils.INTEGER `lb:0,ub:63`
	BSrs utils.INTEGER `lb:0,ub:3`
	BHop utils.INTEGER `lb:0,ub:3`
}
