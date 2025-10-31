package ies

import "rrc/utils"

// SRS-CC-SetIndex ::= SEQUENCE
type SrsCcSetindex struct {
	CcSetindex        *utils.INTEGER `lb:0,ub:3`
	CcIndexinoneccSet *utils.INTEGER `lb:0,ub:7`
}
