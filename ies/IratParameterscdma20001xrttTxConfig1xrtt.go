package ies

import "rrc/utils"

// IRAT-ParametersCDMA2000-1XRTT-tx-Config1XRTT ::= ENUMERATED
type IratParameterscdma20001xrttTxConfig1xrtt struct {
	Value utils.ENUMERATED
}

const (
	IratParameterscdma20001xrttTxConfig1xrttEnumeratedNothing = iota
	IratParameterscdma20001xrttTxConfig1xrttEnumeratedSingle
	IratParameterscdma20001xrttTxConfig1xrttEnumeratedDual
)
