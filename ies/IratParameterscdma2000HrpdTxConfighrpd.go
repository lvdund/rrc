package ies

import "rrc/utils"

// IRAT-ParametersCDMA2000-HRPD-tx-ConfigHRPD ::= ENUMERATED
type IratParameterscdma2000HrpdTxConfighrpd struct {
	Value utils.ENUMERATED
}

const (
	IratParameterscdma2000HrpdTxConfighrpdEnumeratedNothing = iota
	IratParameterscdma2000HrpdTxConfighrpdEnumeratedSingle
	IratParameterscdma2000HrpdTxConfighrpdEnumeratedDual
)
