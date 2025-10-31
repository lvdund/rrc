package ies

import "rrc/utils"

// MeasScaleFactor-r12 ::= ENUMERATED
type MeasscalefactorR12 struct {
	Value utils.ENUMERATED
}

const (
	MeasscalefactorR12EnumeratedNothing = iota
	MeasscalefactorR12EnumeratedSf_Eutra_Cf1
	MeasscalefactorR12EnumeratedSf_Eutra_Cf2
)
