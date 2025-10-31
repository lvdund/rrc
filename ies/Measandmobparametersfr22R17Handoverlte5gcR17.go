package ies

import "rrc/utils"

// MeasAndMobParametersFR2-2-r17-handoverLTE-5GC-r17 ::= ENUMERATED
type Measandmobparametersfr22R17Handoverlte5gcR17 struct {
	Value utils.ENUMERATED
}

const (
	Measandmobparametersfr22R17Handoverlte5gcR17EnumeratedNothing = iota
	Measandmobparametersfr22R17Handoverlte5gcR17EnumeratedSupported
)
