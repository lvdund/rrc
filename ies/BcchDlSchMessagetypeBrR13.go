package ies

import "rrc/utils"

// BCCH-DL-SCH-MessageType-BR-r13 ::= CHOICE
type BcchDlSchMessagetypeBrR13 interface {
	isBcchDlSchMessagetypeBrR13()
}

type BcchDlSchMessagetypeBrR13C1 struct {
	Value interface{}
}

func (*BcchDlSchMessagetypeBrR13C1) isBcchDlSchMessagetypeBrR13() {}

type BcchDlSchMessagetypeBrR13Messageclassextension struct {
	Value interface{}
}

func (*BcchDlSchMessagetypeBrR13Messageclassextension) isBcchDlSchMessagetypeBrR13() {}
