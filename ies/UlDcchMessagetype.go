package ies

import "rrc/utils"

// UL-DCCH-MessageType ::= CHOICE
type UlDcchMessagetype interface {
	isUlDcchMessagetype()
}

type UlDcchMessagetypeC1 struct {
	Value interface{}
}

func (*UlDcchMessagetypeC1) isUlDcchMessagetype() {}

type UlDcchMessagetypeMessageclassextension struct {
	Value interface{}
}

func (*UlDcchMessagetypeMessageclassextension) isUlDcchMessagetype() {}
