package ies

import "rrc/utils"

// TDD-UL-DL-AlignmentOffset-NB-r15 ::= ENUMERATED
type TddUlDlAlignmentoffsetNbR15 struct {
	Value utils.ENUMERATED
}

const (
	TddUlDlAlignmentoffsetNbR15Khz7dot5 = 0
	TddUlDlAlignmentoffsetNbR15Khz0     = 1
	TddUlDlAlignmentoffsetNbR15Khz7dot5 = 2
)
