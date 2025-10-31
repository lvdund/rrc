package ies

import "rrc/utils"

// PDCP-ParametersNR-r15-sn-SizeLo-r15 ::= ENUMERATED
type PdcpParametersnrR15SnSizeloR15 struct {
	Value utils.ENUMERATED
}

const (
	PdcpParametersnrR15SnSizeloR15EnumeratedNothing = iota
	PdcpParametersnrR15SnSizeloR15EnumeratedSupported
)
