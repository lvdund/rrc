package ies

import "rrc/utils"

// CA-ParametersNR-v1730-mux-HARQ-ACK-UnicastMulticast-r17 ::= ENUMERATED
type CaParametersnrV1730MuxHarqAckUnicastmulticastR17 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1730MuxHarqAckUnicastmulticastR17EnumeratedNothing = iota
	CaParametersnrV1730MuxHarqAckUnicastmulticastR17EnumeratedSupported
)
