package ies

import "rrc/utils"

// CA-ParametersNR-v1730-mode2-TDM-CodebookForMux-UnicastMulticastHARQ-ACK-r17 ::= ENUMERATED
type CaParametersnrV1730Mode2TdmCodebookformuxUnicastmulticastharqAckR17 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1730Mode2TdmCodebookformuxUnicastmulticastharqAckR17EnumeratedNothing = iota
	CaParametersnrV1730Mode2TdmCodebookformuxUnicastmulticastharqAckR17EnumeratedSupported
)
