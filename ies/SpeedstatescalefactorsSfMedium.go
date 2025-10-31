package ies

import "rrc/utils"

// SpeedStateScaleFactors-sf-Medium ::= ENUMERATED
type SpeedstatescalefactorsSfMedium struct {
	Value utils.ENUMERATED
}

const (
	SpeedstatescalefactorsSfMediumEnumeratedNothing = iota
	SpeedstatescalefactorsSfMediumEnumeratedOdot25
	SpeedstatescalefactorsSfMediumEnumeratedOdot5
	SpeedstatescalefactorsSfMediumEnumeratedOdot75
	SpeedstatescalefactorsSfMediumEnumeratedLdot0
)
