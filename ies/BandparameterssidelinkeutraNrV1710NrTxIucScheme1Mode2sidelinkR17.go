package ies

import "rrc/utils"

// BandParametersSidelinkEUTRA-NR-v1710-nr-tx-IUC-Scheme1-Mode2Sidelink-r17 ::= ENUMERATED
type BandparameterssidelinkeutraNrV1710NrTxIucScheme1Mode2sidelinkR17 struct {
	Value utils.ENUMERATED
}

const (
	BandparameterssidelinkeutraNrV1710NrTxIucScheme1Mode2sidelinkR17EnumeratedNothing = iota
	BandparameterssidelinkeutraNrV1710NrTxIucScheme1Mode2sidelinkR17EnumeratedSupported
)
