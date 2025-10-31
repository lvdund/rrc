package ies

import "rrc/utils"

// SpeedStateScaleFactors-sf-High ::= ENUMERATED
type SpeedstatescalefactorsSfHigh struct {
	Value utils.ENUMERATED
}

const (
	SpeedstatescalefactorsSfHighEnumeratedNothing = iota
	SpeedstatescalefactorsSfHighEnumeratedOdot25
	SpeedstatescalefactorsSfHighEnumeratedOdot5
	SpeedstatescalefactorsSfHighEnumeratedOdot75
	SpeedstatescalefactorsSfHighEnumeratedLdot0
)
