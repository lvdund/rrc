package ies

import "rrc/utils"

// IRAT-ParametersCDMA2000-1XRTT-v920 ::= SEQUENCE
type IratParameterscdma20001xrttV920 struct {
	ECsfb1xrttR9          utils.ENUMERATED
	ECsfbConcpsMob1xrttR9 *utils.ENUMERATED
}
