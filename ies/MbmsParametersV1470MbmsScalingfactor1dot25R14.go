package ies

import "rrc/utils"

// MBMS-Parameters-v1470-mbms-ScalingFactor1dot25-r14 ::= ENUMERATED
type MbmsParametersV1470MbmsScalingfactor1dot25R14 struct {
	Value utils.ENUMERATED
}

const (
	MbmsParametersV1470MbmsScalingfactor1dot25R14EnumeratedNothing = iota
	MbmsParametersV1470MbmsScalingfactor1dot25R14EnumeratedN3
	MbmsParametersV1470MbmsScalingfactor1dot25R14EnumeratedN6
	MbmsParametersV1470MbmsScalingfactor1dot25R14EnumeratedN9
	MbmsParametersV1470MbmsScalingfactor1dot25R14EnumeratedN12
)
