package ies

import "rrc/utils"

// TDD-UL-DL-AlignmentOffset-NB-r15 ::= ENUMERATED
type TddUlDlAlignmentoffsetNbR15 struct {
	Value utils.ENUMERATED
}

const (
	TddUlDlAlignmentoffsetNbR15EnumeratedNothing = iota
	TddUlDlAlignmentoffsetNbR15EnumeratedKhz_7dot5
	TddUlDlAlignmentoffsetNbR15EnumeratedKhz0
	TddUlDlAlignmentoffsetNbR15EnumeratedKhz7dot5
)
