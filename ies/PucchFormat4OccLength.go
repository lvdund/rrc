package ies

import "rrc/utils"

// PUCCH-format4-occ-Length ::= ENUMERATED
type PucchFormat4OccLength struct {
	Value utils.ENUMERATED
}

const (
	PucchFormat4OccLengthEnumeratedNothing = iota
	PucchFormat4OccLengthEnumeratedN2
	PucchFormat4OccLengthEnumeratedN4
)
