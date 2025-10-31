package ies

import "rrc/utils"

// Other-Parameters-v1430-rlm-ReportSupport-r14 ::= ENUMERATED
type OtherParametersV1430RlmReportsupportR14 struct {
	Value utils.ENUMERATED
}

const (
	OtherParametersV1430RlmReportsupportR14EnumeratedNothing = iota
	OtherParametersV1430RlmReportsupportR14EnumeratedSupported
)
