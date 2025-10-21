package ies

import "rrc/utils"

// PhyLayerParameters-v1130 ::= SEQUENCE
type PhylayerparametersV1130 struct {
	CrsInterfhandlR11       *utils.ENUMERATED
	EpdcchR11               *utils.ENUMERATED
	MultiackCsiReportingR11 *utils.ENUMERATED
	SsCchInterfhandlR11     *utils.ENUMERATED
	TddSpecialsubframeR11   *utils.ENUMERATED
	TxdivPucch1bChselectR11 *utils.ENUMERATED
	UlCompR11               *utils.ENUMERATED
}
