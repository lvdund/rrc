package ies

import "rrc/utils"

// CarrierState-r17 ::= CHOICE
const (
	CarrierstateR17ChoiceNothing = iota
	CarrierstateR17ChoiceDeactivatedR17
	CarrierstateR17ChoiceActivebwpR17
)

type CarrierstateR17 struct {
	Choice         uint64
	DeactivatedR17 *struct{}
	ActivebwpR17   *utils.INTEGER `lb:0,ub:maxNrofBWPs`
}
