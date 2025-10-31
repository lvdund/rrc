package ies

import "rrc/utils"

// MRDC-Parameters-v1620-tdm-restrictionTDD-endc-r16 ::= ENUMERATED
type MrdcParametersV1620TdmRestrictiontddEndcR16 struct {
	Value utils.ENUMERATED
}

const (
	MrdcParametersV1620TdmRestrictiontddEndcR16EnumeratedNothing = iota
	MrdcParametersV1620TdmRestrictiontddEndcR16EnumeratedSupported
)
