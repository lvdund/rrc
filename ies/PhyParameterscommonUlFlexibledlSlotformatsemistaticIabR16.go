package ies

import "rrc/utils"

// Phy-ParametersCommon-ul-flexibleDL-SlotFormatSemiStatic-IAB-r16 ::= ENUMERATED
type PhyParameterscommonUlFlexibledlSlotformatsemistaticIabR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonUlFlexibledlSlotformatsemistaticIabR16EnumeratedNothing = iota
	PhyParameterscommonUlFlexibledlSlotformatsemistaticIabR16EnumeratedSupported
)
