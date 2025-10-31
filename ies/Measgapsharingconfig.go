package ies

import "rrc/utils"

// MeasGapSharingConfig ::= SEQUENCE
// Extensible
type Measgapsharingconfig struct {
	Gapsharingfr2 *utils.Setuprelease[Measgapsharingscheme]
	Gapsharingfr1 *utils.Setuprelease[Measgapsharingscheme] `ext`
	Gapsharingue  *utils.Setuprelease[Measgapsharingscheme] `ext`
}
