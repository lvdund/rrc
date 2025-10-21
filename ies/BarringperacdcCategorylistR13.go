package ies

import "rrc/utils"

// BarringPerACDC-CategoryList-r13 ::= SEQUENCE OF BarringPerACDC-Category-r13
// SIZE (1..maxACDC-Cat-r13)
type BarringperacdcCategorylistR13 struct {
	Value []BarringperacdcCategoryR13
}
