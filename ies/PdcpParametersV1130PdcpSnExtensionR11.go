package ies

import "rrc/utils"

// PDCP-Parameters-v1130-pdcp-SN-Extension-r11 ::= ENUMERATED
type PdcpParametersV1130PdcpSnExtensionR11 struct {
	Value utils.ENUMERATED
}

const (
	PdcpParametersV1130PdcpSnExtensionR11EnumeratedNothing = iota
	PdcpParametersV1130PdcpSnExtensionR11EnumeratedSupported
)
