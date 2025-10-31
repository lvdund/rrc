package ies

// PhyLayerParameters-v1130 ::= SEQUENCE
type PhylayerparametersV1130 struct {
	CrsInterfhandlR11       *PhylayerparametersV1130CrsInterfhandlR11
	EpdcchR11               *PhylayerparametersV1130EpdcchR11
	MultiackCsiReportingR11 *PhylayerparametersV1130MultiackCsiReportingR11
	SsCchInterfhandlR11     *PhylayerparametersV1130SsCchInterfhandlR11
	TddSpecialsubframeR11   *PhylayerparametersV1130TddSpecialsubframeR11
	TxdivPucch1bChselectR11 *PhylayerparametersV1130TxdivPucch1bChselectR11
	UlCompR11               *PhylayerparametersV1130UlCompR11
}
