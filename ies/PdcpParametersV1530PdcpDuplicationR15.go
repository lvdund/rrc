package ies

import "rrc/utils"

// PDCP-Parameters-v1530-pdcp-Duplication-r15 ::= ENUMERATED
type PdcpParametersV1530PdcpDuplicationR15 struct {
	Value utils.ENUMERATED
}

const (
	PdcpParametersV1530PdcpDuplicationR15EnumeratedNothing = iota
	PdcpParametersV1530PdcpDuplicationR15EnumeratedSupported
)
