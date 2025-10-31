package ies

import "rrc/utils"

// MRDC-Parameters-v1630-interBandMRDC-WithOverlapDL-Bands-r16 ::= ENUMERATED
type MrdcParametersV1630InterbandmrdcWithoverlapdlBandsR16 struct {
	Value utils.ENUMERATED
}

const (
	MrdcParametersV1630InterbandmrdcWithoverlapdlBandsR16EnumeratedNothing = iota
	MrdcParametersV1630InterbandmrdcWithoverlapdlBandsR16EnumeratedSupported
)
