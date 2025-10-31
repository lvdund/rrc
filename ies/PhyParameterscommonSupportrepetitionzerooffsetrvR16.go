package ies

import "rrc/utils"

// Phy-ParametersCommon-supportRepetitionZeroOffsetRV-r16 ::= ENUMERATED
type PhyParameterscommonSupportrepetitionzerooffsetrvR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonSupportrepetitionzerooffsetrvR16EnumeratedNothing = iota
	PhyParameterscommonSupportrepetitionzerooffsetrvR16EnumeratedSupported
)
