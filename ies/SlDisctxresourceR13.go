package ies

import "rrc/utils"

// SL-DiscTxResource-r13 ::= CHOICE
type SlDisctxresourceR13 interface {
	isSlDisctxresourceR13()
}

type SlDisctxresourceR13Release struct {
	Value struct{}
}

func (*SlDisctxresourceR13Release) isSlDisctxresourceR13() {}

type SlDisctxresourceR13Setup struct {
	Value interface{}
}

func (*SlDisctxresourceR13Setup) isSlDisctxresourceR13() {}
