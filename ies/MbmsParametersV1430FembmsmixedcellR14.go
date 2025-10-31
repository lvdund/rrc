package ies

import "rrc/utils"

// MBMS-Parameters-v1430-fembmsMixedCell-r14 ::= ENUMERATED
type MbmsParametersV1430FembmsmixedcellR14 struct {
	Value utils.ENUMERATED
}

const (
	MbmsParametersV1430FembmsmixedcellR14EnumeratedNothing = iota
	MbmsParametersV1430FembmsmixedcellR14EnumeratedSupported
)
