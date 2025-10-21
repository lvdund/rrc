package ies

import "rrc/utils"

// SRS-CC-SetIndex-r14 ::= SEQUENCE
type SrsCcSetindexR14 struct {
	CcSetindexR14        utils.INTEGER
	CcIndexinoneccSetR14 utils.INTEGER
}
