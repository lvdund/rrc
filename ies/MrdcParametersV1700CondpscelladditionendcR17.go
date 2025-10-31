package ies

import "rrc/utils"

// MRDC-Parameters-v1700-condPSCellAdditionENDC-r17 ::= ENUMERATED
type MrdcParametersV1700CondpscelladditionendcR17 struct {
	Value utils.ENUMERATED
}

const (
	MrdcParametersV1700CondpscelladditionendcR17EnumeratedNothing = iota
	MrdcParametersV1700CondpscelladditionendcR17EnumeratedSupported
)
