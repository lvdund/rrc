package ies

import "rrc/utils"

// TDD-PUSCH-UpPTS-r14 ::= CHOICE
type TddPuschUpptsR14 interface {
	isTddPuschUpptsR14()
}

type TddPuschUpptsR14Release struct {
	Value struct{}
}

func (*TddPuschUpptsR14Release) isTddPuschUpptsR14() {}

type TddPuschUpptsR14Setup struct {
	Value interface{}
}

func (*TddPuschUpptsR14Setup) isTddPuschUpptsR14() {}
