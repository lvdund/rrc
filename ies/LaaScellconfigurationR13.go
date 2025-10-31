package ies

import "rrc/utils"

// LAA-SCellConfiguration-r13 ::= SEQUENCE
type LaaScellconfigurationR13 struct {
	SubframestartpositionR13  LaaScellconfigurationR13SubframestartpositionR13
	LaaScellsubframeconfigR13 utils.BITSTRING `lb:8,ub:8`
}
