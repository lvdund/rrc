package ies

import "rrc/utils"

// Phy-ParametersCommon-t-DeltaReceptionSupport-IAB-r16 ::= ENUMERATED
type PhyParameterscommonTDeltareceptionsupportIabR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonTDeltareceptionsupportIabR16EnumeratedNothing = iota
	PhyParameterscommonTDeltareceptionsupportIabR16EnumeratedSupported
)
