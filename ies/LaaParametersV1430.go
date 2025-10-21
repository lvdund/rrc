package ies

import "rrc/utils"

// LAA-Parameters-v1430 ::= SEQUENCE
type LaaParametersV1430 struct {
	CrosscarrierschedulinglaaUlR14 *utils.ENUMERATED
	UplinklaaR14                   *utils.ENUMERATED
	TwostepschedulingtiminginfoR14 *utils.ENUMERATED
	UssBlinddecodingadjustmentR14  *utils.ENUMERATED
	UssBlinddecodingreductionR14   *utils.ENUMERATED
	OutofsequencegranthandlingR14  *utils.ENUMERATED
}
