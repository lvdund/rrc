package ies

import "rrc/utils"

// DL-PPW-PreConfig-r17 ::= SEQUENCE
type DlPpwPreconfigR17 struct {
	DlPpwIdR17                      DlPpwIdR17
	DlPpwPeriodicityandstartslotR17 DlPpwPeriodicityandstartslotR17
	LengthR17                       utils.INTEGER `lb:0,ub:160`
	TypeR17                         *DlPpwPreconfigR17TypeR17
	PriorityR17                     *DlPpwPreconfigR17PriorityR17
}
