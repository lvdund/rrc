package ies

import "rrc/utils"

// CA-ParametersNR-v1730-fdm-CodebookForMux-UnicastMulticastHARQ-ACK-r17 ::= ENUMERATED
type CaParametersnrV1730FdmCodebookformuxUnicastmulticastharqAckR17 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1730FdmCodebookformuxUnicastmulticastharqAckR17EnumeratedNothing = iota
	CaParametersnrV1730FdmCodebookformuxUnicastmulticastharqAckR17EnumeratedSupported
)
