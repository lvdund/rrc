package ies

import "rrc/utils"

// SL-TF-IndexPair-r12 ::= SEQUENCE
type SlTfIndexpairR12 struct {
	DiscsfIndexR12  *utils.INTEGER `lb:0,ub:200`
	DiscprbIndexR12 *utils.INTEGER `lb:0,ub:50`
}
