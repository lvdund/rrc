package ies

import "rrc/utils"

// CA-ParametersNRDC-v1610-intraFR-NR-DC-PwrSharingMode2-r16 ::= ENUMERATED
type CaParametersnrdcV1610IntrafrNrDcPwrsharingmode2R16 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrdcV1610IntrafrNrDcPwrsharingmode2R16EnumeratedNothing = iota
	CaParametersnrdcV1610IntrafrNrDcPwrsharingmode2R16EnumeratedSupported
)
