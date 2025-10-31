package ies

import "rrc/utils"

// SL-PSBCH-Config-r16 ::= SEQUENCE
// Extensible
type SlPsbchConfigR16 struct {
	DlP0PsbchR16    *utils.INTEGER `lb:0,ub:15`
	DlAlphaPsbchR16 *SlPsbchConfigR16DlAlphaPsbchR16
	DlP0PsbchR17    *utils.INTEGER `lb:0,ub:24,ext`
}
