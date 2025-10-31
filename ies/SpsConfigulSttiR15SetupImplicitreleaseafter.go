package ies

import "rrc/utils"

// SPS-ConfigUL-STTI-r15-setup-implicitReleaseAfter ::= ENUMERATED
type SpsConfigulSttiR15SetupImplicitreleaseafter struct {
	Value utils.ENUMERATED
}

const (
	SpsConfigulSttiR15SetupImplicitreleaseafterEnumeratedNothing = iota
	SpsConfigulSttiR15SetupImplicitreleaseafterEnumeratedE2
	SpsConfigulSttiR15SetupImplicitreleaseafterEnumeratedE3
	SpsConfigulSttiR15SetupImplicitreleaseafterEnumeratedE4
	SpsConfigulSttiR15SetupImplicitreleaseafterEnumeratedE8
)
