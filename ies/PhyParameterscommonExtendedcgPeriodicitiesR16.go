package ies

import "rrc/utils"

// Phy-ParametersCommon-extendedCG-Periodicities-r16 ::= ENUMERATED
type PhyParameterscommonExtendedcgPeriodicitiesR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonExtendedcgPeriodicitiesR16EnumeratedNothing = iota
	PhyParameterscommonExtendedcgPeriodicitiesR16EnumeratedSupported
)
