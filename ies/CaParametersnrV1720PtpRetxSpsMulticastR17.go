package ies

import "rrc/utils"

// CA-ParametersNR-v1720-ptp-Retx-SPS-Multicast-r17 ::= ENUMERATED
type CaParametersnrV1720PtpRetxSpsMulticastR17 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1720PtpRetxSpsMulticastR17EnumeratedNothing = iota
	CaParametersnrV1720PtpRetxSpsMulticastR17EnumeratedSupported
)
