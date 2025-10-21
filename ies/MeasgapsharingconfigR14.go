package ies

import "rrc/utils"

// MeasGapSharingConfig-r14 ::= CHOICE
type MeasgapsharingconfigR14 interface {
	isMeasgapsharingconfigR14()
}

type MeasgapsharingconfigR14Release struct {
	Value struct{}
}

func (*MeasgapsharingconfigR14Release) isMeasgapsharingconfigR14() {}

type MeasgapsharingconfigR14Setup struct {
	Value interface{}
}

func (*MeasgapsharingconfigR14Setup) isMeasgapsharingconfigR14() {}
