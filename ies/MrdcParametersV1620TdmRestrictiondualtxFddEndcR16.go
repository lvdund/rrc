package ies

import "rrc/utils"

// MRDC-Parameters-v1620-tdm-restrictionDualTX-FDD-endc-r16 ::= ENUMERATED
type MrdcParametersV1620TdmRestrictiondualtxFddEndcR16 struct {
	Value utils.ENUMERATED
}

const (
	MrdcParametersV1620TdmRestrictiondualtxFddEndcR16EnumeratedNothing = iota
	MrdcParametersV1620TdmRestrictiondualtxFddEndcR16EnumeratedSupported
)
