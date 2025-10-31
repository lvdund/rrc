package ies

import "rrc/utils"

// PDCP-Parameters-v1130-supportRohcContextContinue-r11 ::= ENUMERATED
type PdcpParametersV1130SupportrohccontextcontinueR11 struct {
	Value utils.ENUMERATED
}

const (
	PdcpParametersV1130SupportrohccontextcontinueR11EnumeratedNothing = iota
	PdcpParametersV1130SupportrohccontextcontinueR11EnumeratedSupported
)
