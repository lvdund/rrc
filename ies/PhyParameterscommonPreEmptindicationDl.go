package ies

import "rrc/utils"

// Phy-ParametersCommon-pre-EmptIndication-DL ::= ENUMERATED
type PhyParameterscommonPreEmptindicationDl struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonPreEmptindicationDlEnumeratedNothing = iota
	PhyParameterscommonPreEmptindicationDlEnumeratedSupported
)
