package ies

import "rrc/utils"

// PCCH-MessageType ::= CHOICE
type PcchMessagetype interface {
	isPcchMessagetype()
}

type PcchMessagetypeC1 struct {
	Value interface{}
}

func (*PcchMessagetypeC1) isPcchMessagetype() {}

type PcchMessagetypeMessageclassextension struct {
	Value interface{}
}

func (*PcchMessagetypeMessageclassextension) isPcchMessagetype() {}
