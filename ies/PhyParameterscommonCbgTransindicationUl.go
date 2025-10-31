package ies

import "rrc/utils"

// Phy-ParametersCommon-cbg-TransIndication-UL ::= ENUMERATED
type PhyParameterscommonCbgTransindicationUl struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonCbgTransindicationUlEnumeratedNothing = iota
	PhyParameterscommonCbgTransindicationUlEnumeratedSupported
)
