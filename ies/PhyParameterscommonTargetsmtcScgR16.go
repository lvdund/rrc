package ies

import "rrc/utils"

// Phy-ParametersCommon-targetSMTC-SCG-r16 ::= ENUMERATED
type PhyParameterscommonTargetsmtcScgR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonTargetsmtcScgR16EnumeratedNothing = iota
	PhyParameterscommonTargetsmtcScgR16EnumeratedSupported
)
