package ies

import "rrc/utils"

// CA-ParametersNR-v1700-disablingScalingFactorDormantSCell-r17 ::= ENUMERATED
type CaParametersnrV1700DisablingscalingfactordormantscellR17 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1700DisablingscalingfactordormantscellR17EnumeratedNothing = iota
	CaParametersnrV1700DisablingscalingfactordormantscellR17EnumeratedSupported
)
