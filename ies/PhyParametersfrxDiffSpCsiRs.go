package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-sp-CSI-RS ::= ENUMERATED
type PhyParametersfrxDiffSpCsiRs struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffSpCsiRsEnumeratedNothing = iota
	PhyParametersfrxDiffSpCsiRsEnumeratedSupported
)
