package ies

import "rrc/utils"

// MeasGapConfig ::= CHOICE
// Extensible
type Measgapconfig interface {
	isMeasgapconfig()
}

type MeasgapconfigRelease struct {
	Value struct{}
}

func (*MeasgapconfigRelease) isMeasgapconfig() {}

type MeasgapconfigSetup struct {
	Value interface{}
}

func (*MeasgapconfigSetup) isMeasgapconfig() {}
