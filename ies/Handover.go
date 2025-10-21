package ies

import "rrc/utils"

// Handover ::= SEQUENCE
// Extensible
type Handover struct {
	TargetratType             utils.ENUMERATED
	TargetratMessagecontainer utils.OCTETSTRING
	NasSecurityparamfromeutra *utils.OCTETSTRING
	Systeminformation         *SiOrpsiGeran
}
