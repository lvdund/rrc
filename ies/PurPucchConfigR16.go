package ies

import "rrc/utils"

// PUR-PUCCH-Config-r16 ::= SEQUENCE
type PurPucchConfigR16 struct {
	N1pucchAnR16                   *utils.INTEGER
	PucchNumrepetitionceFormat1R16 *utils.ENUMERATED
}
