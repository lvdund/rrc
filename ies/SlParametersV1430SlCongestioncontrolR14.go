package ies

import "rrc/utils"

// SL-Parameters-v1430-sl-CongestionControl-r14 ::= ENUMERATED
type SlParametersV1430SlCongestioncontrolR14 struct {
	Value utils.ENUMERATED
}

const (
	SlParametersV1430SlCongestioncontrolR14EnumeratedNothing = iota
	SlParametersV1430SlCongestioncontrolR14EnumeratedSupported
)
