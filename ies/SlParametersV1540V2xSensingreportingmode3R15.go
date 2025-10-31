package ies

import "rrc/utils"

// SL-Parameters-v1540-v2x-SensingReportingMode3-r15 ::= ENUMERATED
type SlParametersV1540V2xSensingreportingmode3R15 struct {
	Value utils.ENUMERATED
}

const (
	SlParametersV1540V2xSensingreportingmode3R15EnumeratedNothing = iota
	SlParametersV1540V2xSensingreportingmode3R15EnumeratedSupported
)
