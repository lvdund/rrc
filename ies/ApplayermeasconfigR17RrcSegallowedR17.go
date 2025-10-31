package ies

import "rrc/utils"

// AppLayerMeasConfig-r17-rrc-SegAllowed-r17 ::= ENUMERATED
type ApplayermeasconfigR17RrcSegallowedR17 struct {
	Value utils.ENUMERATED
}

const (
	ApplayermeasconfigR17RrcSegallowedR17EnumeratedNothing = iota
	ApplayermeasconfigR17RrcSegallowedR17EnumeratedEnabled
)
