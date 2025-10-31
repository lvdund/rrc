package ies

import "rrc/utils"

// MeasAndMobParametersCommon-condHandoverParametersCommon-r16-condHandoverFDD-TDD-r16 ::= ENUMERATED
type MeasandmobparameterscommonCondhandoverparameterscommonR16CondhandoverfddTddR16 struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparameterscommonCondhandoverparameterscommonR16CondhandoverfddTddR16EnumeratedNothing = iota
	MeasandmobparameterscommonCondhandoverparameterscommonR16CondhandoverfddTddR16EnumeratedSupported
)
