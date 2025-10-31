package ies

import "rrc/utils"

// Phy-ParametersCommon-harqACK-jointMultiDCI-MultiTRP-r16 ::= ENUMERATED
type PhyParameterscommonHarqackJointmultidciMultitrpR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonHarqackJointmultidciMultitrpR16EnumeratedNothing = iota
	PhyParameterscommonHarqackJointmultidciMultitrpR16EnumeratedSupported
)
