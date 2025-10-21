package ies

import "rrc/utils"

// LAA-SCellConfiguration-r13 ::= SEQUENCE
type LaaScellconfigurationR13 struct {
	SubframestartpositionR13  utils.ENUMERATED
	LaaScellsubframeconfigR13 utils.BITSTRING
}
