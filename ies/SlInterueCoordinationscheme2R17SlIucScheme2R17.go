package ies

import "rrc/utils"

// SL-InterUE-CoordinationScheme2-r17-sl-IUC-Scheme2-r17 ::= ENUMERATED
type SlInterueCoordinationscheme2R17SlIucScheme2R17 struct {
	Value utils.ENUMERATED
}

const (
	SlInterueCoordinationscheme2R17SlIucScheme2R17EnumeratedNothing = iota
	SlInterueCoordinationscheme2R17SlIucScheme2R17EnumeratedEnabled
)
