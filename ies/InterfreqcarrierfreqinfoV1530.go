package ies

import "rrc/utils"

// InterFreqCarrierFreqInfo-v1530 ::= SEQUENCE
type InterfreqcarrierfreqinfoV1530 struct {
	HsdnIndicationR15             utils.BOOLEAN
	InterfreqneighhsdnCelllistR15 *InterfreqneighhsdnCelllistR15
	CellselectioninfoceV1530      *CellselectioninfoceV1530
}
