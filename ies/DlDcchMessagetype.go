package ies

import "rrc/utils"

// DL-DCCH-MessageType ::= CHOICE
type DlDcchMessagetype interface {
	isDlDcchMessagetype()
}

type DlDcchMessagetypeC1 struct {
	Value interface{}
}

func (*DlDcchMessagetypeC1) isDlDcchMessagetype() {}

type DlDcchMessagetypeMessageclassextension struct {
	Value interface{}
}

func (*DlDcchMessagetypeMessageclassextension) isDlDcchMessagetype() {}
