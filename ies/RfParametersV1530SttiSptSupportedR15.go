package ies

import "rrc/utils"

// RF-Parameters-v1530-sTTI-SPT-Supported-r15 ::= ENUMERATED
type RfParametersV1530SttiSptSupportedR15 struct {
	Value utils.ENUMERATED
}

const (
	RfParametersV1530SttiSptSupportedR15EnumeratedNothing = iota
	RfParametersV1530SttiSptSupportedR15EnumeratedSupported
)
