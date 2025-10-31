package ies

import "rrc/utils"

// MeasAndMobParametersFR2-2-r17-idleInactiveNR-MeasReport-r17 ::= ENUMERATED
type Measandmobparametersfr22R17IdleinactivenrMeasreportR17 struct {
	Value utils.ENUMERATED
}

const (
	Measandmobparametersfr22R17IdleinactivenrMeasreportR17EnumeratedNothing = iota
	Measandmobparametersfr22R17IdleinactivenrMeasreportR17EnumeratedSupported
)
