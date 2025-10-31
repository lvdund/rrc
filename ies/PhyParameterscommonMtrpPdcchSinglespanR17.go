package ies

import "rrc/utils"

// Phy-ParametersCommon-mTRP-PDCCH-singleSpan-r17 ::= ENUMERATED
type PhyParameterscommonMtrpPdcchSinglespanR17 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonMtrpPdcchSinglespanR17EnumeratedNothing = iota
	PhyParameterscommonMtrpPdcchSinglespanR17EnumeratedSupported
)
