package ies

import "rrc/utils"

// CA-ParametersNR-v1730-pucch-ConfigForSPS-Multicast-r17 ::= ENUMERATED
type CaParametersnrV1730PucchConfigforspsMulticastR17 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1730PucchConfigforspsMulticastR17EnumeratedNothing = iota
	CaParametersnrV1730PucchConfigforspsMulticastR17EnumeratedSupported
)
