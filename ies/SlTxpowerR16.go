package ies

import "rrc/utils"

// SL-TxPower-r16 ::= CHOICE
const (
	SlTxpowerR16ChoiceNothing = iota
	SlTxpowerR16ChoiceMinusinfinityR16
	SlTxpowerR16ChoiceTxpowerR16
)

type SlTxpowerR16 struct {
	Choice           uint64
	MinusinfinityR16 *struct{}
	TxpowerR16       *utils.INTEGER `lb:-30,ub:33`
}
