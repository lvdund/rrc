package ies

import "rrc/utils"

// EDT-TBS-NB-r15-edt-TBS-r15 ::= ENUMERATED
type EdtTbsNbR15EdtTbsR15 struct {
	Value utils.ENUMERATED
}

const (
	EdtTbsNbR15EdtTbsR15EnumeratedNothing = iota
	EdtTbsNbR15EdtTbsR15EnumeratedB328
	EdtTbsNbR15EdtTbsR15EnumeratedB408
	EdtTbsNbR15EdtTbsR15EnumeratedB504
	EdtTbsNbR15EdtTbsR15EnumeratedB584
	EdtTbsNbR15EdtTbsR15EnumeratedB680
	EdtTbsNbR15EdtTbsR15EnumeratedB808
	EdtTbsNbR15EdtTbsR15EnumeratedB936
	EdtTbsNbR15EdtTbsR15EnumeratedB1000
)
