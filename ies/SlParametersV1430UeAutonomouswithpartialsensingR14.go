package ies

import "rrc/utils"

// SL-Parameters-v1430-ue-AutonomousWithPartialSensing-r14 ::= ENUMERATED
type SlParametersV1430UeAutonomouswithpartialsensingR14 struct {
	Value utils.ENUMERATED
}

const (
	SlParametersV1430UeAutonomouswithpartialsensingR14EnumeratedNothing = iota
	SlParametersV1430UeAutonomouswithpartialsensingR14EnumeratedSupported
)
