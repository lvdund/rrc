package ies

import "rrc/utils"

// BandParametersSidelink-v1710-tx-IUC-Scheme1-Mode2Sidelink-r17 ::= ENUMERATED
type BandparameterssidelinkV1710TxIucScheme1Mode2sidelinkR17 struct {
	Value utils.ENUMERATED
}

const (
	BandparameterssidelinkV1710TxIucScheme1Mode2sidelinkR17EnumeratedNothing = iota
	BandparameterssidelinkV1710TxIucScheme1Mode2sidelinkR17EnumeratedSupported
)
