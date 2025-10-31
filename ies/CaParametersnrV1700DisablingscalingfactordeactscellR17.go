package ies

import "rrc/utils"

// CA-ParametersNR-v1700-disablingScalingFactorDeactSCell-r17 ::= ENUMERATED
type CaParametersnrV1700DisablingscalingfactordeactscellR17 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1700DisablingscalingfactordeactscellR17EnumeratedNothing = iota
	CaParametersnrV1700DisablingscalingfactordeactscellR17EnumeratedSupported
)
