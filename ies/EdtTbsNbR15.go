package ies

import "rrc/utils"

// EDT-TBS-NB-r15 ::= SEQUENCE
type EdtTbsNbR15 struct {
	EdtSmalltbsEnabledR15 utils.BOOLEAN
	EdtTbsR15             EdtTbsNbR15EdtTbsR15
}
