package ies

import "rrc/utils"

// UAC-BarringPerCat-NB-r16 ::= SEQUENCE
type UacBarringpercatNbR16 struct {
	UacAccesscategoryR16 utils.INTEGER `lb:0,ub:maxAccessCat1R15`
	UacBarringfactorR16  UacBarringpercatNbR16UacBarringfactorR16
	UacBarringtimeR16    UacBarringpercatNbR16UacBarringtimeR16
}
