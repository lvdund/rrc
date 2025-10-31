package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-uci-CodeBlockSegmentation ::= ENUMERATED
type PhyParametersfrxDiffUciCodeblocksegmentation struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffUciCodeblocksegmentationEnumeratedNothing = iota
	PhyParametersfrxDiffUciCodeblocksegmentationEnumeratedSupported
)
