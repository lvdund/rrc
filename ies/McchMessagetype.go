package ies

import "rrc/utils"

// MCCH-MessageType ::= CHOICE
type McchMessagetype interface {
	isMcchMessagetype()
}

type McchMessagetypeC1 struct {
	Value interface{}
}

func (*McchMessagetypeC1) isMcchMessagetype() {}

type McchMessagetypeLater struct {
	Value interface{}
}

func (*McchMessagetypeLater) isMcchMessagetype() {}
