package ies

import "rrc/utils"

// MBMS-Parameters-v1430-subcarrierSpacingMBMS-khz7dot5-r14 ::= ENUMERATED
type MbmsParametersV1430SubcarrierspacingmbmsKhz7dot5R14 struct {
	Value utils.ENUMERATED
}

const (
	MbmsParametersV1430SubcarrierspacingmbmsKhz7dot5R14EnumeratedNothing = iota
	MbmsParametersV1430SubcarrierspacingmbmsKhz7dot5R14EnumeratedSupported
)
