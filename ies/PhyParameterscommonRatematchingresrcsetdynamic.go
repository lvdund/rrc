package ies

import "rrc/utils"

// Phy-ParametersCommon-rateMatchingResrcSetDynamic ::= ENUMERATED
type PhyParameterscommonRatematchingresrcsetdynamic struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonRatematchingresrcsetdynamicEnumeratedNothing = iota
	PhyParameterscommonRatematchingresrcsetdynamicEnumeratedSupported
)
