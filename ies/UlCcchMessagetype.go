package ies

import "rrc/utils"

// UL-CCCH-MessageType ::= CHOICE
type UlCcchMessagetype interface {
	isUlCcchMessagetype()
}

type UlCcchMessagetypeC1 struct {
	Value interface{}
}

func (*UlCcchMessagetypeC1) isUlCcchMessagetype() {}

type UlCcchMessagetypeMessageclassextension struct {
	Value interface{}
}

func (*UlCcchMessagetypeMessageclassextension) isUlCcchMessagetype() {}
