package ies

import "rrc/utils"

// BandParametersSidelinkEUTRA-NR-v1630-nr-rx-Sidelink-r16 ::= ENUMERATED
type BandparameterssidelinkeutraNrV1630NrRxSidelinkR16 struct {
	Value utils.ENUMERATED
}

const (
	BandparameterssidelinkeutraNrV1630NrRxSidelinkR16EnumeratedNothing = iota
	BandparameterssidelinkeutraNrV1630NrRxSidelinkR16EnumeratedSupported
)
