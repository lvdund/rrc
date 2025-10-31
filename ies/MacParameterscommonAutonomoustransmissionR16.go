package ies

import "rrc/utils"

// MAC-ParametersCommon-autonomousTransmission-r16 ::= ENUMERATED
type MacParameterscommonAutonomoustransmissionR16 struct {
	Value utils.ENUMERATED
}

const (
	MacParameterscommonAutonomoustransmissionR16EnumeratedNothing = iota
	MacParameterscommonAutonomoustransmissionR16EnumeratedSupported
)
