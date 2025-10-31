package ies

import "rrc/utils"

// BarringPerACDC-Category-r13 ::= SEQUENCE
type BarringperacdcCategoryR13 struct {
	AcdcCategoryR13      utils.INTEGER `lb:0,ub:maxACDCCatR13`
	AcdcBarringconfigR13 *BarringperacdcCategoryR13AcdcBarringconfigR13
}
