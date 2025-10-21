package ies

import "rrc/utils"

// MBMS-Parameters-v1470 ::= SEQUENCE
type MbmsParametersV1470 struct {
	MbmsMaxbwR14               interface{}
	MbmsScalingfactor1dot25R14 *utils.ENUMERATED
	MbmsScalingfactor7dot5R14  *utils.ENUMERATED
}
