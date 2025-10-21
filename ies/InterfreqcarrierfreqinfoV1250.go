package ies

import "rrc/utils"

// InterFreqCarrierFreqInfo-v1250 ::= SEQUENCE
type InterfreqcarrierfreqinfoV1250 struct {
	ReducedmeasperformanceR12   *utils.ENUMERATED
	QQualminrsrqOnallsymbolsR12 *QQualminR9
}
