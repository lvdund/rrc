package ies

// IRAT-ParametersCDMA2000-HRPD ::= SEQUENCE
type IratParameterscdma2000Hrpd struct {
	Supportedbandlisthrpd Supportedbandlisthrpd
	TxConfighrpd          IratParameterscdma2000HrpdTxConfighrpd
	RxConfighrpd          IratParameterscdma2000HrpdRxConfighrpd
}
