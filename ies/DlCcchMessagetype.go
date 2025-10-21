package ies

import "rrc/utils"

// DL-CCCH-MessageType ::= CHOICE
type DlCcchMessagetype interface {
	isDlCcchMessagetype()
}

type DlCcchMessagetypeC1 struct {
	Value interface{}
}

func (*DlCcchMessagetypeC1) isDlCcchMessagetype() {}

type DlCcchMessagetypeMessageclassextension struct {
	Value interface{}
}

func (*DlCcchMessagetypeMessageclassextension) isDlCcchMessagetype() {}
