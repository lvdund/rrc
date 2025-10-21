package ies

import "rrc/utils"

// TPC-Index ::= CHOICE
type TpcIndex interface {
	isTpcIndex()
}

type TpcIndexIndexofformat3 struct {
	Value utils.INTEGER
}

func (*TpcIndexIndexofformat3) isTpcIndex() {}

type TpcIndexIndexofformat3a struct {
	Value utils.INTEGER
}

func (*TpcIndexIndexofformat3a) isTpcIndex() {}
