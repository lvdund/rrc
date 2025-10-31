package ies

import "rrc/utils"

// Phy-ParametersCommon-separateRACH-IAB-Support-r16 ::= ENUMERATED
type PhyParameterscommonSeparaterachIabSupportR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonSeparaterachIabSupportR16EnumeratedNothing = iota
	PhyParameterscommonSeparaterachIabSupportR16EnumeratedSupported
)
