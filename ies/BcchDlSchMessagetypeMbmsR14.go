package ies

import "rrc/utils"

// BCCH-DL-SCH-MessageType-MBMS-r14 ::= CHOICE
type BcchDlSchMessagetypeMbmsR14 interface {
	isBcchDlSchMessagetypeMbmsR14()
}

type BcchDlSchMessagetypeMbmsR14C1 struct {
	Value interface{}
}

func (*BcchDlSchMessagetypeMbmsR14C1) isBcchDlSchMessagetypeMbmsR14() {}

type BcchDlSchMessagetypeMbmsR14Messageclassextension struct {
	Value interface{}
}

func (*BcchDlSchMessagetypeMbmsR14Messageclassextension) isBcchDlSchMessagetypeMbmsR14() {}
