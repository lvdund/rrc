package ies

import "rrc/utils"

// Phy-ParametersCommon-extendedSPS-Periodicities-r16 ::= ENUMERATED
type PhyParameterscommonExtendedspsPeriodicitiesR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonExtendedspsPeriodicitiesR16EnumeratedNothing = iota
	PhyParameterscommonExtendedspsPeriodicitiesR16EnumeratedSupported
)
