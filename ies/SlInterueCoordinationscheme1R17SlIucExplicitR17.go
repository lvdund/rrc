package ies

import "rrc/utils"

// SL-InterUE-CoordinationScheme1-r17-sl-IUC-Explicit-r17 ::= ENUMERATED
type SlInterueCoordinationscheme1R17SlIucExplicitR17 struct {
	Value utils.ENUMERATED
}

const (
	SlInterueCoordinationscheme1R17SlIucExplicitR17EnumeratedNothing = iota
	SlInterueCoordinationscheme1R17SlIucExplicitR17EnumeratedEnabled
	SlInterueCoordinationscheme1R17SlIucExplicitR17EnumeratedDisabled
)
