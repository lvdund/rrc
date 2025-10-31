package ies

import "rrc/utils"

// SL-InterUE-CoordinationScheme1-r17-sl-Condition1-A-2-r17 ::= ENUMERATED
type SlInterueCoordinationscheme1R17SlCondition1A2R17 struct {
	Value utils.ENUMERATED
}

const (
	SlInterueCoordinationscheme1R17SlCondition1A2R17EnumeratedNothing = iota
	SlInterueCoordinationscheme1R17SlCondition1A2R17EnumeratedDisabled
)
