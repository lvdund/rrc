package ies

import "rrc/utils"

// PUR-PUCCH-Config-r16 ::= SEQUENCE
type PurPucchConfigR16 struct {
	N1pucchAnR16                   *utils.INTEGER `lb:0,ub:2047`
	PucchNumrepetitionceFormat1R16 *PurPucchConfigR16PucchNumrepetitionceFormat1R16
}
