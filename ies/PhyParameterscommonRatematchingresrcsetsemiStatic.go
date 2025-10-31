package ies

import "rrc/utils"

// Phy-ParametersCommon-rateMatchingResrcSetSemi-Static ::= ENUMERATED
type PhyParameterscommonRatematchingresrcsetsemiStatic struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonRatematchingresrcsetsemiStaticEnumeratedNothing = iota
	PhyParameterscommonRatematchingresrcsetsemiStaticEnumeratedSupported
)
