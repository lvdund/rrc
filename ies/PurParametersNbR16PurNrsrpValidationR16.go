package ies

import "rrc/utils"

// PUR-Parameters-NB-r16-pur-NRSRP-Validation-r16 ::= ENUMERATED
type PurParametersNbR16PurNrsrpValidationR16 struct {
	Value utils.ENUMERATED
}

const (
	PurParametersNbR16PurNrsrpValidationR16EnumeratedNothing = iota
	PurParametersNbR16PurNrsrpValidationR16EnumeratedSupported
)
