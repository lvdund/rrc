package ies

// IRAT-ParametersCDMA2000-1XRTT ::= SEQUENCE
type IratParameterscdma20001xrtt struct {
	Supportedbandlist1xrtt Supportedbandlist1xrtt
	TxConfig1xrtt          IratParameterscdma20001xrttTxConfig1xrtt
	RxConfig1xrtt          IratParameterscdma20001xrttRxConfig1xrtt
}
