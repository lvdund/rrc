package ies

// CA-ParametersNRDC-v1610 ::= SEQUENCE
type CaParametersnrdcV1610 struct {
	IntrafrNrDcPwrsharingmode1R16   *CaParametersnrdcV1610IntrafrNrDcPwrsharingmode1R16
	IntrafrNrDcPwrsharingmode2R16   *CaParametersnrdcV1610IntrafrNrDcPwrsharingmode2R16
	IntrafrNrDcDynamicpwrsharingR16 *CaParametersnrdcV1610IntrafrNrDcDynamicpwrsharingR16
	AsyncnrdcR16                    *CaParametersnrdcV1610AsyncnrdcR16
}
