package ies

import "rrc/utils"

// BCCH-DL-SCH-MessageType-NB ::= CHOICE
type BcchDlSchMessagetypeNb interface {
	isBcchDlSchMessagetypeNb()
}

type BcchDlSchMessagetypeNbC1 struct {
	Value interface{}
}

func (*BcchDlSchMessagetypeNbC1) isBcchDlSchMessagetypeNb() {}

type BcchDlSchMessagetypeNbMessageclassextension struct {
	Value interface{}
}

func (*BcchDlSchMessagetypeNbMessageclassextension) isBcchDlSchMessagetypeNb() {}
