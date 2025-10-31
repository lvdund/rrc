package ies

import "rrc/utils"

// MRDC-Parameters-v1620-tdm-restrictionFDD-endc-r16 ::= ENUMERATED
type MrdcParametersV1620TdmRestrictionfddEndcR16 struct {
	Value utils.ENUMERATED
}

const (
	MrdcParametersV1620TdmRestrictionfddEndcR16EnumeratedNothing = iota
	MrdcParametersV1620TdmRestrictionfddEndcR16EnumeratedSupported
)
