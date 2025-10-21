package ies

import "rrc/utils"

// MeasScaleFactor-r12 ::= ENUMERATED
type MeasscalefactorR12 struct {
	Value utils.ENUMERATED
}

const (
	MeasscalefactorR12SfEutraCf1 = 0
	MeasscalefactorR12SfEutraCf2 = 1
)
