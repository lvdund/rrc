package ies

import "rrc/utils"

// BandParametersSidelinkEUTRA-NR-v1630-nr-sl-CrossCarrierScheduling-r16 ::= ENUMERATED
type BandparameterssidelinkeutraNrV1630NrSlCrosscarrierschedulingR16 struct {
	Value utils.ENUMERATED
}

const (
	BandparameterssidelinkeutraNrV1630NrSlCrosscarrierschedulingR16EnumeratedNothing = iota
	BandparameterssidelinkeutraNrV1630NrSlCrosscarrierschedulingR16EnumeratedSupported
)
