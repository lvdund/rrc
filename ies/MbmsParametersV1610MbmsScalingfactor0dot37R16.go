package ies

import "rrc/utils"

// MBMS-Parameters-v1610-mbms-ScalingFactor0dot37-r16 ::= ENUMERATED
type MbmsParametersV1610MbmsScalingfactor0dot37R16 struct {
	Value utils.ENUMERATED
}

const (
	MbmsParametersV1610MbmsScalingfactor0dot37R16EnumeratedNothing = iota
	MbmsParametersV1610MbmsScalingfactor0dot37R16EnumeratedN12
	MbmsParametersV1610MbmsScalingfactor0dot37R16EnumeratedN16
	MbmsParametersV1610MbmsScalingfactor0dot37R16EnumeratedN20
	MbmsParametersV1610MbmsScalingfactor0dot37R16EnumeratedN24
)
