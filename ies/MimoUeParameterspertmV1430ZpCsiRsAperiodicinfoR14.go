package ies

import "rrc/utils"

// MIMO-UE-ParametersPerTM-v1430-zp-CSI-RS-AperiodicInfo-r14 ::= ENUMERATED
type MimoUeParameterspertmV1430ZpCsiRsAperiodicinfoR14 struct {
	Value utils.ENUMERATED
}

const (
	MimoUeParameterspertmV1430ZpCsiRsAperiodicinfoR14EnumeratedNothing = iota
	MimoUeParameterspertmV1430ZpCsiRsAperiodicinfoR14EnumeratedSupported
)
