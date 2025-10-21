package ies

import "rrc/utils"

// MBMS-Parameters-v1610 ::= SEQUENCE
type MbmsParametersV1610 struct {
	MbmsScalingfactor2dot5R16    *utils.ENUMERATED
	MbmsScalingfactor0dot37R16   *utils.ENUMERATED
	MbmsSupportedbandinfolistR16 interface{}
}
