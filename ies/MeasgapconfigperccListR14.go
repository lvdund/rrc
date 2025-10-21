package ies

import "rrc/utils"

// MeasGapConfigPerCC-List-r14 ::= CHOICE
type MeasgapconfigperccListR14 interface {
	isMeasgapconfigperccListR14()
}

type MeasgapconfigperccListR14Release struct {
	Value struct{}
}

func (*MeasgapconfigperccListR14Release) isMeasgapconfigperccListR14() {}

type MeasgapconfigperccListR14Setup struct {
	Value interface{}
}

func (*MeasgapconfigperccListR14Setup) isMeasgapconfigperccListR14() {}
