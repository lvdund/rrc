package ies

import "rrc/utils"

// Phy-ParametersCommon-dci-DL-PriorityIndicator-r16 ::= ENUMERATED
type PhyParameterscommonDciDlPriorityindicatorR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonDciDlPriorityindicatorR16EnumeratedNothing = iota
	PhyParameterscommonDciDlPriorityindicatorR16EnumeratedSupported
)
