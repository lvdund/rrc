package ies

import "rrc/utils"

// UE-RadioPagingInfo-NB-r13-ue-Category-NB-r13 ::= ENUMERATED
type UeRadiopaginginfoNbR13UeCategoryNbR13 struct {
	Value utils.ENUMERATED
}

const (
	UeRadiopaginginfoNbR13UeCategoryNbR13EnumeratedNothing = iota
	UeRadiopaginginfoNbR13UeCategoryNbR13EnumeratedNb1
)
