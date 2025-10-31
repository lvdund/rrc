package ies

import "rrc/utils"

// SL-InterUE-CoordinationScheme2-r17-sl-IndicationUE-B-r17 ::= ENUMERATED
type SlInterueCoordinationscheme2R17SlIndicationueBR17 struct {
	Value utils.ENUMERATED
}

const (
	SlInterueCoordinationscheme2R17SlIndicationueBR17EnumeratedNothing = iota
	SlInterueCoordinationscheme2R17SlIndicationueBR17EnumeratedEnabled
	SlInterueCoordinationscheme2R17SlIndicationueBR17EnumeratedDisabled
)
