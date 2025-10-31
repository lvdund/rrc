package ies

import "rrc/utils"

// Phy-ParametersCommon-cbg-FlushIndication-DL ::= ENUMERATED
type PhyParameterscommonCbgFlushindicationDl struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonCbgFlushindicationDlEnumeratedNothing = iota
	PhyParameterscommonCbgFlushindicationDlEnumeratedSupported
)
