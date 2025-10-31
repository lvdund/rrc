package ies

import "rrc/utils"

// SL-Parameters-v1430-slss-TxRx-r14 ::= ENUMERATED
type SlParametersV1430SlssTxrxR14 struct {
	Value utils.ENUMERATED
}

const (
	SlParametersV1430SlssTxrxR14EnumeratedNothing = iota
	SlParametersV1430SlssTxrxR14EnumeratedSupported
)
