package ies

import "rrc/utils"

// UAC-BarringPerCat-r15 ::= SEQUENCE
type UacBarringpercatR15 struct {
	AccesscategoryR15         utils.INTEGER `lb:0,ub:maxAccessCat1R15`
	UacBarringinfosetindexR15 UacBarringinfosetindexR15
}
