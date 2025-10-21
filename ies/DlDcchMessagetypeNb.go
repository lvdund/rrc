package ies

import "rrc/utils"

// DL-DCCH-MessageType-NB ::= CHOICE
type DlDcchMessagetypeNb interface {
	isDlDcchMessagetypeNb()
}

type DlDcchMessagetypeNbC1 struct {
	Value interface{}
}

func (*DlDcchMessagetypeNbC1) isDlDcchMessagetypeNb() {}

type DlDcchMessagetypeNbMessageclassextension struct {
	Value interface{}
}

func (*DlDcchMessagetypeNbMessageclassextension) isDlDcchMessagetypeNb() {}
