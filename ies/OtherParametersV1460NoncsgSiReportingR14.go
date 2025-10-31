package ies

import "rrc/utils"

// Other-Parameters-v1460-nonCSG-SI-Reporting-r14 ::= ENUMERATED
type OtherParametersV1460NoncsgSiReportingR14 struct {
	Value utils.ENUMERATED
}

const (
	OtherParametersV1460NoncsgSiReportingR14EnumeratedNothing = iota
	OtherParametersV1460NoncsgSiReportingR14EnumeratedSupported
)
