package ies

import "rrc/utils"

// MeasAndMobParametersCommon-idleInactiveEUTRA-MeasReport-r16 ::= ENUMERATED
type MeasandmobparameterscommonIdleinactiveeutraMeasreportR16 struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparameterscommonIdleinactiveeutraMeasreportR16EnumeratedNothing = iota
	MeasandmobparameterscommonIdleinactiveeutraMeasreportR16EnumeratedSupported
)
