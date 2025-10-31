package ies

import "rrc/utils"

// Phy-ParametersCommon-separateSMTC-InterIAB-Support-r16 ::= ENUMERATED
type PhyParameterscommonSeparatesmtcInteriabSupportR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonSeparatesmtcInteriabSupportR16EnumeratedNothing = iota
	PhyParameterscommonSeparatesmtcInteriabSupportR16EnumeratedSupported
)
