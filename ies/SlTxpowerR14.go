package ies

import "rrc/utils"

// SL-TxPower-r14 ::= CHOICE
const (
	SlTxpowerR14ChoiceNothing = iota
	SlTxpowerR14ChoiceMinusinfinityR14
	SlTxpowerR14ChoiceTxpowerR14
)

type SlTxpowerR14 struct {
	Choice           uint64
	MinusinfinityR14 *struct{}
	TxpowerR14       *utils.INTEGER `lb:-41,ub:31`
}
