package ies

import "rrc/utils"

// PUCCH-format4-occ-Index ::= ENUMERATED
type PucchFormat4OccIndex struct {
	Value utils.ENUMERATED
}

const (
	PucchFormat4OccIndexEnumeratedNothing = iota
	PucchFormat4OccIndexEnumeratedN0
	PucchFormat4OccIndexEnumeratedN1
	PucchFormat4OccIndexEnumeratedN2
	PucchFormat4OccIndexEnumeratedN3
)
