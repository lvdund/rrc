package ies

import "rrc/utils"

// AB-Config-NB-r13-ab-Category-r13 ::= ENUMERATED
type AbConfigNbR13AbCategoryR13 struct {
	Value utils.ENUMERATED
}

const (
	AbConfigNbR13AbCategoryR13EnumeratedNothing = iota
	AbConfigNbR13AbCategoryR13EnumeratedA
	AbConfigNbR13AbCategoryR13EnumeratedB
	AbConfigNbR13AbCategoryR13EnumeratedC
)
