package ies

import "rrc/utils"

// SL-Parameters-v1540 ::= SEQUENCE
type SlParametersV1540 struct {
	Sl64qamRxR15                *utils.ENUMERATED
	SlRatematchingtbsscalingR15 *utils.ENUMERATED
	SlLowt2minR15               *utils.ENUMERATED
	V2xSensingreportingmode3R15 *utils.ENUMERATED
}
