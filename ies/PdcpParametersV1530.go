package ies

// PDCP-Parameters-v1530 ::= SEQUENCE
type PdcpParametersV1530 struct {
	SupportedudcR15    *SupportedudcR15
	PdcpDuplicationR15 *PdcpParametersV1530PdcpDuplicationR15
}
