package ies

import "rrc/utils"

// UL-CCCH-MessageType-NB ::= CHOICE
type UlCcchMessagetypeNb interface {
	isUlCcchMessagetypeNb()
}

type UlCcchMessagetypeNbC1 struct {
	Value interface{}
}

func (*UlCcchMessagetypeNbC1) isUlCcchMessagetypeNb() {}

type UlCcchMessagetypeNbMessageclassextension struct {
	Value interface{}
}

func (*UlCcchMessagetypeNbMessageclassextension) isUlCcchMessagetypeNb() {}
