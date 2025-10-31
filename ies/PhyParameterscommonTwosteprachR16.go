package ies

import "rrc/utils"

// Phy-ParametersCommon-twoStepRACH-r16 ::= ENUMERATED
type PhyParameterscommonTwosteprachR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonTwosteprachR16EnumeratedNothing = iota
	PhyParameterscommonTwosteprachR16EnumeratedSupported
)
