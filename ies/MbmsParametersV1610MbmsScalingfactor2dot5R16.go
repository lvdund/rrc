package ies

import "rrc/utils"

// MBMS-Parameters-v1610-mbms-ScalingFactor2dot5-r16 ::= ENUMERATED
type MbmsParametersV1610MbmsScalingfactor2dot5R16 struct {
	Value utils.ENUMERATED
}

const (
	MbmsParametersV1610MbmsScalingfactor2dot5R16EnumeratedNothing = iota
	MbmsParametersV1610MbmsScalingfactor2dot5R16EnumeratedN2
	MbmsParametersV1610MbmsScalingfactor2dot5R16EnumeratedN4
	MbmsParametersV1610MbmsScalingfactor2dot5R16EnumeratedN6
	MbmsParametersV1610MbmsScalingfactor2dot5R16EnumeratedN8
)
