package ies

import "rrc/utils"

// Phy-ParametersCommon-ul-flexibleDL-SlotFormatDynamics-IAB-r16 ::= ENUMERATED
type PhyParameterscommonUlFlexibledlSlotformatdynamicsIabR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonUlFlexibledlSlotformatdynamicsIabR16EnumeratedNothing = iota
	PhyParameterscommonUlFlexibledlSlotformatdynamicsIabR16EnumeratedSupported
)
