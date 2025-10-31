package ies

import "rrc/utils"

// MBMS-Parameters-v1470-mbms-ScalingFactor7dot5-r14 ::= ENUMERATED
type MbmsParametersV1470MbmsScalingfactor7dot5R14 struct {
	Value utils.ENUMERATED
}

const (
	MbmsParametersV1470MbmsScalingfactor7dot5R14EnumeratedNothing = iota
	MbmsParametersV1470MbmsScalingfactor7dot5R14EnumeratedN1
	MbmsParametersV1470MbmsScalingfactor7dot5R14EnumeratedN2
	MbmsParametersV1470MbmsScalingfactor7dot5R14EnumeratedN3
	MbmsParametersV1470MbmsScalingfactor7dot5R14EnumeratedN4
)
