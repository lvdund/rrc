package ies

import "rrc/utils"

// SPS-ConfigUL-STTI-r15 ::= CHOICE
// Extensible
type SpsConfigulSttiR15 interface {
	isSpsConfigulSttiR15()
}

type SpsConfigulSttiR15Release struct {
	Value struct{}
}

func (*SpsConfigulSttiR15Release) isSpsConfigulSttiR15() {}

type SpsConfigulSttiR15Setup struct {
	Value interface{}
}

func (*SpsConfigulSttiR15Setup) isSpsConfigulSttiR15() {}
