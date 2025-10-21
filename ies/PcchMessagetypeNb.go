package ies

import "rrc/utils"

// PCCH-MessageType-NB ::= CHOICE
type PcchMessagetypeNb interface {
	isPcchMessagetypeNb()
}

type PcchMessagetypeNbC1 struct {
	Value interface{}
}

func (*PcchMessagetypeNbC1) isPcchMessagetypeNb() {}

type PcchMessagetypeNbMessageclassextension struct {
	Value interface{}
}

func (*PcchMessagetypeNbMessageclassextension) isPcchMessagetypeNb() {}
