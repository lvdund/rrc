package ies

import "rrc/utils"

// MeasAndMobParametersFR2-2-r17-handoverInterF-r17 ::= ENUMERATED
type Measandmobparametersfr22R17HandoverinterfR17 struct {
	Value utils.ENUMERATED
}

const (
	Measandmobparametersfr22R17HandoverinterfR17EnumeratedNothing = iota
	Measandmobparametersfr22R17HandoverinterfR17EnumeratedSupported
)
