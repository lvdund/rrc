package ies

import "rrc/utils"

// CA-ParametersNR-v1640-uplinkTxDC-TwoCarrierReport-r16 ::= ENUMERATED
type CaParametersnrV1640UplinktxdcTwocarrierreportR16 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1640UplinktxdcTwocarrierreportR16EnumeratedNothing = iota
	CaParametersnrV1640UplinktxdcTwocarrierreportR16EnumeratedSupported
)
