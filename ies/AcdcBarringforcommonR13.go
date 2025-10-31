package ies

import "rrc/utils"

// ACDC-BarringForCommon-r13 ::= SEQUENCE
type AcdcBarringforcommonR13 struct {
	AcdcHplmnonlyR13              utils.BOOLEAN
	BarringperacdcCategorylistR13 BarringperacdcCategorylistR13
}
