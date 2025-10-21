package ies

import "rrc/utils"

// TDM-PatternConfig-r15 ::= CHOICE
type TdmPatternconfigR15 interface {
	isTdmPatternconfigR15()
}

type TdmPatternconfigR15Release struct {
	Value struct{}
}

func (*TdmPatternconfigR15Release) isTdmPatternconfigR15() {}

type TdmPatternconfigR15Setup struct {
	Value interface{}
}

func (*TdmPatternconfigR15Setup) isTdmPatternconfigR15() {}
