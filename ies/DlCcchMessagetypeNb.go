package ies

import "rrc/utils"

// DL-CCCH-MessageType-NB ::= CHOICE
type DlCcchMessagetypeNb interface {
	isDlCcchMessagetypeNb()
}

type DlCcchMessagetypeNbC1 struct {
	Value interface{}
}

func (*DlCcchMessagetypeNbC1) isDlCcchMessagetypeNb() {}

type DlCcchMessagetypeNbMessageclassextension struct {
	Value interface{}
}

func (*DlCcchMessagetypeNbMessageclassextension) isDlCcchMessagetypeNb() {}
