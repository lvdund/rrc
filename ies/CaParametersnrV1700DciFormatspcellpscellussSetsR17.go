package ies

import "rrc/utils"

// CA-ParametersNR-v1700-dci-FormatsPCellPSCellUSS-Sets-r17 ::= ENUMERATED
type CaParametersnrV1700DciFormatspcellpscellussSetsR17 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1700DciFormatspcellpscellussSetsR17EnumeratedNothing = iota
	CaParametersnrV1700DciFormatspcellpscellussSetsR17EnumeratedSupported
)
