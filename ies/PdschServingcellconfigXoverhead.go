package ies

import "rrc/utils"

// PDSCH-ServingCellConfig-xOverhead ::= ENUMERATED
type PdschServingcellconfigXoverhead struct {
	Value utils.ENUMERATED
}

const (
	PdschServingcellconfigXoverheadEnumeratedNothing = iota
	PdschServingcellconfigXoverheadEnumeratedXoh6
	PdschServingcellconfigXoverheadEnumeratedXoh12
	PdschServingcellconfigXoverheadEnumeratedXoh18
)
