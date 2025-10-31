package ies

import "rrc/utils"

// GapConfig-r17-gapType-r17 ::= ENUMERATED
type GapconfigR17GaptypeR17 struct {
	Value utils.ENUMERATED
}

const (
	GapconfigR17GaptypeR17EnumeratedNothing = iota
	GapconfigR17GaptypeR17EnumeratedPerue
	GapconfigR17GaptypeR17EnumeratedPerfr1
	GapconfigR17GaptypeR17EnumeratedPerfr2
)
