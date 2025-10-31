package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-aggregationFactorSPS-DL-r16 ::= ENUMERATED
type PhyParametersfrxDiffAggregationfactorspsDlR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffAggregationfactorspsDlR16EnumeratedNothing = iota
	PhyParametersfrxDiffAggregationfactorspsDlR16EnumeratedSupported
)
