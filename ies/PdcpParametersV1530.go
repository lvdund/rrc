package ies

import "rrc/utils"

// PDCP-Parameters-v1530 ::= SEQUENCE
type PdcpParametersV1530 struct {
	SupportedudcR15    *SupportedudcR15
	PdcpDuplicationR15 *utils.ENUMERATED
}
