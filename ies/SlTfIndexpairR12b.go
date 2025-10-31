package ies

import "rrc/utils"

// SL-TF-IndexPair-r12b ::= SEQUENCE
type SlTfIndexpairR12b struct {
	DiscsfIndexR12b  *utils.INTEGER `lb:0,ub:209`
	DiscprbIndexR12b *utils.INTEGER `lb:0,ub:49`
}
