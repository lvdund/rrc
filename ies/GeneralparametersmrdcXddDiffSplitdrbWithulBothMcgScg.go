package ies

import "rrc/utils"

// GeneralParametersMRDC-XDD-Diff-splitDRB-withUL-Both-MCG-SCG ::= ENUMERATED
type GeneralparametersmrdcXddDiffSplitdrbWithulBothMcgScg struct {
	Value utils.ENUMERATED
}

const (
	GeneralparametersmrdcXddDiffSplitdrbWithulBothMcgScgEnumeratedNothing = iota
	GeneralparametersmrdcXddDiffSplitdrbWithulBothMcgScgEnumeratedSupported
)
