package ies

import "rrc/utils"

// BandParametersSidelinkEUTRA-NR-v1630-nr-tx-Sidelink-r16 ::= ENUMERATED
type BandparameterssidelinkeutraNrV1630NrTxSidelinkR16 struct {
	Value utils.ENUMERATED
}

const (
	BandparameterssidelinkeutraNrV1630NrTxSidelinkR16EnumeratedNothing = iota
	BandparameterssidelinkeutraNrV1630NrTxSidelinkR16EnumeratedSupported
)
