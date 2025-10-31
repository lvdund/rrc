package ies

import "rrc/utils"

// Phy-ParametersCommon-rateMatchingCtrlResrcSetDynamic ::= ENUMERATED
type PhyParameterscommonRatematchingctrlresrcsetdynamic struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonRatematchingctrlresrcsetdynamicEnumeratedNothing = iota
	PhyParameterscommonRatematchingctrlresrcsetdynamicEnumeratedSupported
)
