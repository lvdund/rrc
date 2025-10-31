package ies

import "rrc/utils"

// IRAT-ParametersCDMA2000-1XRTT-rx-Config1XRTT ::= ENUMERATED
type IratParameterscdma20001xrttRxConfig1xrtt struct {
	Value utils.ENUMERATED
}

const (
	IratParameterscdma20001xrttRxConfig1xrttEnumeratedNothing = iota
	IratParameterscdma20001xrttRxConfig1xrttEnumeratedSingle
	IratParameterscdma20001xrttRxConfig1xrttEnumeratedDual
)
