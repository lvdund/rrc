package ies

import "rrc/utils"

// SL-Parameters-v1430-ue-AutonomousWithFullSensing-r14 ::= ENUMERATED
type SlParametersV1430UeAutonomouswithfullsensingR14 struct {
	Value utils.ENUMERATED
}

const (
	SlParametersV1430UeAutonomouswithfullsensingR14EnumeratedNothing = iota
	SlParametersV1430UeAutonomouswithfullsensingR14EnumeratedSupported
)
