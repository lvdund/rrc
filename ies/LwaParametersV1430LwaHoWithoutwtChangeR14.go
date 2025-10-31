package ies

import "rrc/utils"

// LWA-Parameters-v1430-lwa-HO-WithoutWT-Change-r14 ::= ENUMERATED
type LwaParametersV1430LwaHoWithoutwtChangeR14 struct {
	Value utils.ENUMERATED
}

const (
	LwaParametersV1430LwaHoWithoutwtChangeR14EnumeratedNothing = iota
	LwaParametersV1430LwaHoWithoutwtChangeR14EnumeratedSupported
)
