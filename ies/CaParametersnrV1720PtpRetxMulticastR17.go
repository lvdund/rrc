package ies

import "rrc/utils"

// CA-ParametersNR-v1720-ptp-Retx-Multicast-r17 ::= ENUMERATED
type CaParametersnrV1720PtpRetxMulticastR17 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1720PtpRetxMulticastR17EnumeratedNothing = iota
	CaParametersnrV1720PtpRetxMulticastR17EnumeratedSupported
)
