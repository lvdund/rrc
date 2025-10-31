package ies

import "rrc/utils"

// SON-Parameters-r9-rach-Report-r9 ::= ENUMERATED
type SonParametersR9RachReportR9 struct {
	Value utils.ENUMERATED
}

const (
	SonParametersR9RachReportR9EnumeratedNothing = iota
	SonParametersR9RachReportR9EnumeratedSupported
)
