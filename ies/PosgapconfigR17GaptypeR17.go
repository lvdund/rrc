package ies

import "rrc/utils"

// PosGapConfig-r17-gapType-r17 ::= ENUMERATED
type PosgapconfigR17GaptypeR17 struct {
	Value utils.ENUMERATED
}

const (
	PosgapconfigR17GaptypeR17EnumeratedNothing = iota
	PosgapconfigR17GaptypeR17EnumeratedPerue
	PosgapconfigR17GaptypeR17EnumeratedPerfr1
	PosgapconfigR17GaptypeR17EnumeratedPerfr2
)
