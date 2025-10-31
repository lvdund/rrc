package ies

import "rrc/utils"

// CA-ParametersNRDC-v1610-intraFR-NR-DC-DynamicPwrSharing-r16 ::= ENUMERATED
type CaParametersnrdcV1610IntrafrNrDcDynamicpwrsharingR16 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrdcV1610IntrafrNrDcDynamicpwrsharingR16EnumeratedNothing = iota
	CaParametersnrdcV1610IntrafrNrDcDynamicpwrsharingR16EnumeratedShort
	CaParametersnrdcV1610IntrafrNrDcDynamicpwrsharingR16EnumeratedLong
)
