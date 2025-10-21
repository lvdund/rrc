package ies

import "rrc/utils"

// UL-DCCH-MessageType-NB ::= CHOICE
type UlDcchMessagetypeNb interface {
	isUlDcchMessagetypeNb()
}

type UlDcchMessagetypeNbC1 struct {
	Value interface{}
}

func (*UlDcchMessagetypeNbC1) isUlDcchMessagetypeNb() {}

type UlDcchMessagetypeNbMessageclassextension struct {
	Value interface{}
}

func (*UlDcchMessagetypeNbMessageclassextension) isUlDcchMessagetypeNb() {}
