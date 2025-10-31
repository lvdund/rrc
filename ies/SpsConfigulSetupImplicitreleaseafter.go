package ies

import "rrc/utils"

// SPS-ConfigUL-setup-implicitReleaseAfter ::= ENUMERATED
type SpsConfigulSetupImplicitreleaseafter struct {
	Value utils.ENUMERATED
}

const (
	SpsConfigulSetupImplicitreleaseafterEnumeratedNothing = iota
	SpsConfigulSetupImplicitreleaseafterEnumeratedE2
	SpsConfigulSetupImplicitreleaseafterEnumeratedE3
	SpsConfigulSetupImplicitreleaseafterEnumeratedE4
	SpsConfigulSetupImplicitreleaseafterEnumeratedE8
)
