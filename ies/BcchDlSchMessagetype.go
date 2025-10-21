package ies

import "rrc/utils"

// BCCH-DL-SCH-MessageType ::= CHOICE
type BcchDlSchMessagetype interface {
	isBcchDlSchMessagetype()
}

type BcchDlSchMessagetypeC1 struct {
	Value interface{}
}

func (*BcchDlSchMessagetypeC1) isBcchDlSchMessagetype() {}

type BcchDlSchMessagetypeMessageclassextension struct {
	Value interface{}
}

func (*BcchDlSchMessagetypeMessageclassextension) isBcchDlSchMessagetype() {}
