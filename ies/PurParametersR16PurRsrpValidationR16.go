package ies

import "rrc/utils"

// PUR-Parameters-r16-pur-RSRP-Validation-r16 ::= ENUMERATED
type PurParametersR16PurRsrpValidationR16 struct {
	Value utils.ENUMERATED
}

const (
	PurParametersR16PurRsrpValidationR16EnumeratedNothing = iota
	PurParametersR16PurRsrpValidationR16EnumeratedSupported
)
