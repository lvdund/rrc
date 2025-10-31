package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-cqi-4-BitsSubbandTN-NonSharedSpectrumChAccess-r17 ::= ENUMERATED
type PhyParametersfrxDiffCqi4BitssubbandtnNonsharedspectrumchaccessR17 struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffCqi4BitssubbandtnNonsharedspectrumchaccessR17EnumeratedNothing = iota
	PhyParametersfrxDiffCqi4BitssubbandtnNonsharedspectrumchaccessR17EnumeratedSupported
)
