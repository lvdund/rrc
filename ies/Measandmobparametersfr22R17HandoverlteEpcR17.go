package ies

import "rrc/utils"

// MeasAndMobParametersFR2-2-r17-handoverLTE-EPC-r17 ::= ENUMERATED
type Measandmobparametersfr22R17HandoverlteEpcR17 struct {
	Value utils.ENUMERATED
}

const (
	Measandmobparametersfr22R17HandoverlteEpcR17EnumeratedNothing = iota
	Measandmobparametersfr22R17HandoverlteEpcR17EnumeratedSupported
)
