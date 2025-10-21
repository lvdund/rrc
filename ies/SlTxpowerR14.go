package ies

import "rrc/utils"

// SL-TxPower-r14 ::= CHOICE
type SlTxpowerR14 interface {
	isSlTxpowerR14()
}

type SlTxpowerR14MinusinfinityR14 struct {
	Value struct{}
}

func (*SlTxpowerR14MinusinfinityR14) isSlTxpowerR14() {}

type SlTxpowerR14TxpowerR14 struct {
	Value utils.INTEGER
}

func (*SlTxpowerR14TxpowerR14) isSlTxpowerR14() {}
