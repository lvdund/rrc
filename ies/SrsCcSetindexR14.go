package ies

import "rrc/utils"

// SRS-CC-SetIndex-r14 ::= SEQUENCE
type SrsCcSetindexR14 struct {
	CcSetindexR14        utils.INTEGER `lb:0,ub:3`
	CcIndexinoneccSetR14 utils.INTEGER `lb:0,ub:7`
}
