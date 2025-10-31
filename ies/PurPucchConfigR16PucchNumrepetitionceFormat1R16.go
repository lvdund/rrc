package ies

import "rrc/utils"

// PUR-PUCCH-Config-r16-pucch-NumRepetitionCE-Format1-r16 ::= ENUMERATED
type PurPucchConfigR16PucchNumrepetitionceFormat1R16 struct {
	Value utils.ENUMERATED
}

const (
	PurPucchConfigR16PucchNumrepetitionceFormat1R16EnumeratedNothing = iota
	PurPucchConfigR16PucchNumrepetitionceFormat1R16EnumeratedN1
	PurPucchConfigR16PucchNumrepetitionceFormat1R16EnumeratedN2
	PurPucchConfigR16PucchNumrepetitionceFormat1R16EnumeratedN4
	PurPucchConfigR16PucchNumrepetitionceFormat1R16EnumeratedN8
)
