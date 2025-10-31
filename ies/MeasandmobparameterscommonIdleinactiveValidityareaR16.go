package ies

import "rrc/utils"

// MeasAndMobParametersCommon-idleInactive-ValidityArea-r16 ::= ENUMERATED
type MeasandmobparameterscommonIdleinactiveValidityareaR16 struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparameterscommonIdleinactiveValidityareaR16EnumeratedNothing = iota
	MeasandmobparameterscommonIdleinactiveValidityareaR16EnumeratedSupported
)
