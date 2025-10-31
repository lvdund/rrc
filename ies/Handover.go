package ies

import "rrc/utils"

// Handover ::= SEQUENCE
// Extensible
type Handover struct {
	TargetratType             HandoverTargetratType
	TargetratMessagecontainer utils.OCTETSTRING
	NasSecurityparamfromeutra *utils.OCTETSTRING `lb:1,ub:1`
	Systeminformation         *SiOrpsiGeran
}
