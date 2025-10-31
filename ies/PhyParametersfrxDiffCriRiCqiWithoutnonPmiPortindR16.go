package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-cri-RI-CQI-WithoutNon-PMI-PortInd-r16 ::= ENUMERATED
type PhyParametersfrxDiffCriRiCqiWithoutnonPmiPortindR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffCriRiCqiWithoutnonPmiPortindR16EnumeratedNothing = iota
	PhyParametersfrxDiffCriRiCqiWithoutnonPmiPortindR16EnumeratedSupported
)
