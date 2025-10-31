package ies

import "rrc/utils"

// Phy-ParametersCommon-dci-UL-PriorityIndicator-r16 ::= ENUMERATED
type PhyParameterscommonDciUlPriorityindicatorR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonDciUlPriorityindicatorR16EnumeratedNothing = iota
	PhyParameterscommonDciUlPriorityindicatorR16EnumeratedSupported
)
