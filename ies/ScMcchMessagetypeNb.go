package ies

import "rrc/utils"

// SC-MCCH-MessageType-NB ::= CHOICE
type ScMcchMessagetypeNb interface {
	isScMcchMessagetypeNb()
}

type ScMcchMessagetypeNbC1 struct {
	Value interface{}
}

func (*ScMcchMessagetypeNbC1) isScMcchMessagetypeNb() {}

type ScMcchMessagetypeNbMessageclassextension struct {
	Value interface{}
}

func (*ScMcchMessagetypeNbMessageclassextension) isScMcchMessagetypeNb() {}
