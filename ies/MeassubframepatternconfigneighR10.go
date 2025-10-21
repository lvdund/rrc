package ies

import "rrc/utils"

// MeasSubframePatternConfigNeigh-r10 ::= CHOICE
type MeassubframepatternconfigneighR10 interface {
	isMeassubframepatternconfigneighR10()
}

type MeassubframepatternconfigneighR10Release struct {
	Value struct{}
}

func (*MeassubframepatternconfigneighR10Release) isMeassubframepatternconfigneighR10() {}

type MeassubframepatternconfigneighR10Setup struct {
	Value interface{}
}

func (*MeassubframepatternconfigneighR10Setup) isMeassubframepatternconfigneighR10() {}
