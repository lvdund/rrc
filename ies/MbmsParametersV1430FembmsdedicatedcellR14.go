package ies

import "rrc/utils"

// MBMS-Parameters-v1430-fembmsDedicatedCell-r14 ::= ENUMERATED
type MbmsParametersV1430FembmsdedicatedcellR14 struct {
	Value utils.ENUMERATED
}

const (
	MbmsParametersV1430FembmsdedicatedcellR14EnumeratedNothing = iota
	MbmsParametersV1430FembmsdedicatedcellR14EnumeratedSupported
)
