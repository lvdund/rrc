package ies

// LAA-Parameters-v1430 ::= SEQUENCE
type LaaParametersV1430 struct {
	CrosscarrierschedulinglaaUlR14 *LaaParametersV1430CrosscarrierschedulinglaaUlR14
	UplinklaaR14                   *LaaParametersV1430UplinklaaR14
	TwostepschedulingtiminginfoR14 *LaaParametersV1430TwostepschedulingtiminginfoR14
	UssBlinddecodingadjustmentR14  *LaaParametersV1430UssBlinddecodingadjustmentR14
	UssBlinddecodingreductionR14   *LaaParametersV1430UssBlinddecodingreductionR14
	OutofsequencegranthandlingR14  *LaaParametersV1430OutofsequencegranthandlingR14
}
