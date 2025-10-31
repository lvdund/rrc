package ies

import "rrc/utils"

// T-Reordering ::= ENUMERATED
type TReordering struct {
	Value utils.ENUMERATED
}

const (
	TReorderingEnumeratedNothing = iota
	TReorderingEnumeratedMs0
	TReorderingEnumeratedMs5
	TReorderingEnumeratedMs10
	TReorderingEnumeratedMs15
	TReorderingEnumeratedMs20
	TReorderingEnumeratedMs25
	TReorderingEnumeratedMs30
	TReorderingEnumeratedMs35
	TReorderingEnumeratedMs40
	TReorderingEnumeratedMs45
	TReorderingEnumeratedMs50
	TReorderingEnumeratedMs55
	TReorderingEnumeratedMs60
	TReorderingEnumeratedMs65
	TReorderingEnumeratedMs70
	TReorderingEnumeratedMs75
	TReorderingEnumeratedMs80
	TReorderingEnumeratedMs85
	TReorderingEnumeratedMs90
	TReorderingEnumeratedMs95
	TReorderingEnumeratedMs100
	TReorderingEnumeratedMs110
	TReorderingEnumeratedMs120
	TReorderingEnumeratedMs130
	TReorderingEnumeratedMs140
	TReorderingEnumeratedMs150
	TReorderingEnumeratedMs160
	TReorderingEnumeratedMs170
	TReorderingEnumeratedMs180
	TReorderingEnumeratedMs190
	TReorderingEnumeratedMs200
	TReorderingEnumeratedMs1600_V1310
)
