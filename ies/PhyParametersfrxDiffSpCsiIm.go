package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-sp-CSI-IM ::= ENUMERATED
type PhyParametersfrxDiffSpCsiIm struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffSpCsiImEnumeratedNothing = iota
	PhyParametersfrxDiffSpCsiImEnumeratedSupported
)
