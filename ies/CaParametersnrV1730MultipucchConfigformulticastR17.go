package ies

import "rrc/utils"

// CA-ParametersNR-v1730-multiPUCCH-ConfigForMulticast-r17 ::= ENUMERATED
type CaParametersnrV1730MultipucchConfigformulticastR17 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1730MultipucchConfigformulticastR17EnumeratedNothing = iota
	CaParametersnrV1730MultipucchConfigformulticastR17EnumeratedSupported
)
