package ies

import "rrc/utils"

// UAC-BarringPerCat ::= SEQUENCE
type UacBarringpercat struct {
	Accesscategory         utils.INTEGER `lb:0,ub:maxAccessCat1`
	UacBarringinfosetindex UacBarringinfosetindex
}
