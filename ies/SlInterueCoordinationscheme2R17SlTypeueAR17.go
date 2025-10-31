package ies

import "rrc/utils"

// SL-InterUE-CoordinationScheme2-r17-sl-TypeUE-A-r17 ::= ENUMERATED
type SlInterueCoordinationscheme2R17SlTypeueAR17 struct {
	Value utils.ENUMERATED
}

const (
	SlInterueCoordinationscheme2R17SlTypeueAR17EnumeratedNothing = iota
	SlInterueCoordinationscheme2R17SlTypeueAR17EnumeratedEnabled
)
