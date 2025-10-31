package ies

import "rrc/utils"

// V2X-BandParametersEUTRA-NR-v1630-nr-tx-Sidelink-r16 ::= ENUMERATED
type V2xBandparameterseutraNrV1630NrTxSidelinkR16 struct {
	Value utils.ENUMERATED
}

const (
	V2xBandparameterseutraNrV1630NrTxSidelinkR16EnumeratedNothing = iota
	V2xBandparameterseutraNrV1630NrTxSidelinkR16EnumeratedSupported
)
