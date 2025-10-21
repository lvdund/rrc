package ies

import "rrc/utils"

// SL-DiscTxRefCarrierDedicated-r13 ::= CHOICE
type SlDisctxrefcarrierdedicatedR13 interface {
	isSlDisctxrefcarrierdedicatedR13()
}

type SlDisctxrefcarrierdedicatedR13Pcell struct {
	Value struct{}
}

func (*SlDisctxrefcarrierdedicatedR13Pcell) isSlDisctxrefcarrierdedicatedR13() {}

type SlDisctxrefcarrierdedicatedR13Scell struct {
	Value ScellindexR10
}

func (*SlDisctxrefcarrierdedicatedR13Scell) isSlDisctxrefcarrierdedicatedR13() {}
