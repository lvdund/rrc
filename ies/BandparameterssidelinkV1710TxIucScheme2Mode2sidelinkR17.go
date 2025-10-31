package ies

import "rrc/utils"

// BandParametersSidelink-v1710-tx-IUC-Scheme2-Mode2Sidelink-r17 ::= ENUMERATED
type BandparameterssidelinkV1710TxIucScheme2Mode2sidelinkR17 struct {
	Value utils.ENUMERATED
}

const (
	BandparameterssidelinkV1710TxIucScheme2Mode2sidelinkR17EnumeratedNothing = iota
	BandparameterssidelinkV1710TxIucScheme2Mode2sidelinkR17EnumeratedN4
	BandparameterssidelinkV1710TxIucScheme2Mode2sidelinkR17EnumeratedN8
	BandparameterssidelinkV1710TxIucScheme2Mode2sidelinkR17EnumeratedN16
)
