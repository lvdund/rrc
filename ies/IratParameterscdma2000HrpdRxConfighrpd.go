package ies

import "rrc/utils"

// IRAT-ParametersCDMA2000-HRPD-rx-ConfigHRPD ::= ENUMERATED
type IratParameterscdma2000HrpdRxConfighrpd struct {
	Value utils.ENUMERATED
}

const (
	IratParameterscdma2000HrpdRxConfighrpdEnumeratedNothing = iota
	IratParameterscdma2000HrpdRxConfighrpdEnumeratedSingle
	IratParameterscdma2000HrpdRxConfighrpdEnumeratedDual
)
