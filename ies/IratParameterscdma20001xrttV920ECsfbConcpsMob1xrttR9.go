package ies

import "rrc/utils"

// IRAT-ParametersCDMA2000-1XRTT-v920-e-CSFB-ConcPS-Mob1XRTT-r9 ::= ENUMERATED
type IratParameterscdma20001xrttV920ECsfbConcpsMob1xrttR9 struct {
	Value utils.ENUMERATED
}

const (
	IratParameterscdma20001xrttV920ECsfbConcpsMob1xrttR9EnumeratedNothing = iota
	IratParameterscdma20001xrttV920ECsfbConcpsMob1xrttR9EnumeratedSupported
)
