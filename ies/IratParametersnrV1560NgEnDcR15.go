package ies

import "rrc/utils"

// IRAT-ParametersNR-v1560-ng-EN-DC-r15 ::= ENUMERATED
type IratParametersnrV1560NgEnDcR15 struct {
	Value utils.ENUMERATED
}

const (
	IratParametersnrV1560NgEnDcR15EnumeratedNothing = iota
	IratParametersnrV1560NgEnDcR15EnumeratedSupported
)
