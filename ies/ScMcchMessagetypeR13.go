package ies

import "rrc/utils"

// SC-MCCH-MessageType-r13 ::= CHOICE
type ScMcchMessagetypeR13 interface {
	isScMcchMessagetypeR13()
}

type ScMcchMessagetypeR13C1 struct {
	Value interface{}
}

func (*ScMcchMessagetypeR13C1) isScMcchMessagetypeR13() {}

type ScMcchMessagetypeR13Messageclassextension struct {
	Value interface{}
}

func (*ScMcchMessagetypeR13Messageclassextension) isScMcchMessagetypeR13() {}
