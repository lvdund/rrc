package ies

import "rrc/utils"

// PUSCH-ServingCellConfig-xOverhead ::= ENUMERATED
type PuschServingcellconfigXoverhead struct {
	Value utils.ENUMERATED
}

const (
	PuschServingcellconfigXoverheadEnumeratedNothing = iota
	PuschServingcellconfigXoverheadEnumeratedXoh6
	PuschServingcellconfigXoverheadEnumeratedXoh12
	PuschServingcellconfigXoverheadEnumeratedXoh18
)
