package ies

import "rrc/utils"

// IRAT-ParametersCDMA2000-HRPD ::= SEQUENCE
type IratParameterscdma2000Hrpd struct {
	Supportedbandlisthrpd Supportedbandlisthrpd
	TxConfighrpd          utils.ENUMERATED
	RxConfighrpd          utils.ENUMERATED
}
