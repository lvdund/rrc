package ies

import "rrc/utils"

// LAA-SCellConfiguration-r13-subframeStartPosition-r13 ::= ENUMERATED
type LaaScellconfigurationR13SubframestartpositionR13 struct {
	Value utils.ENUMERATED
}

const (
	LaaScellconfigurationR13SubframestartpositionR13EnumeratedNothing = iota
	LaaScellconfigurationR13SubframestartpositionR13EnumeratedS0
	LaaScellconfigurationR13SubframestartpositionR13EnumeratedS07
)
