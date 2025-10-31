package ies

import "rrc/utils"

// V2X-BandParametersEUTRA-NR-v1630-nr-rx-Sidelink-r16 ::= ENUMERATED
type V2xBandparameterseutraNrV1630NrRxSidelinkR16 struct {
	Value utils.ENUMERATED
}

const (
	V2xBandparameterseutraNrV1630NrRxSidelinkR16EnumeratedNothing = iota
	V2xBandparameterseutraNrV1630NrRxSidelinkR16EnumeratedSupported
)
