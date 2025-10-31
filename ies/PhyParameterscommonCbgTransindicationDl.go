package ies

import "rrc/utils"

// Phy-ParametersCommon-cbg-TransIndication-DL ::= ENUMERATED
type PhyParameterscommonCbgTransindicationDl struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonCbgTransindicationDlEnumeratedNothing = iota
	PhyParameterscommonCbgTransindicationDlEnumeratedSupported
)
