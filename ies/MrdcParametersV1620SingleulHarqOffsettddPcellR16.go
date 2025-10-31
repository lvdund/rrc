package ies

import "rrc/utils"

// MRDC-Parameters-v1620-singleUL-HARQ-offsetTDD-PCell-r16 ::= ENUMERATED
type MrdcParametersV1620SingleulHarqOffsettddPcellR16 struct {
	Value utils.ENUMERATED
}

const (
	MrdcParametersV1620SingleulHarqOffsettddPcellR16EnumeratedNothing = iota
	MrdcParametersV1620SingleulHarqOffsettddPcellR16EnumeratedSupported
)
