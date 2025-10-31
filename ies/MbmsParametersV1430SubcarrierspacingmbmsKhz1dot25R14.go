package ies

import "rrc/utils"

// MBMS-Parameters-v1430-subcarrierSpacingMBMS-khz1dot25-r14 ::= ENUMERATED
type MbmsParametersV1430SubcarrierspacingmbmsKhz1dot25R14 struct {
	Value utils.ENUMERATED
}

const (
	MbmsParametersV1430SubcarrierspacingmbmsKhz1dot25R14EnumeratedNothing = iota
	MbmsParametersV1430SubcarrierspacingmbmsKhz1dot25R14EnumeratedSupported
)
