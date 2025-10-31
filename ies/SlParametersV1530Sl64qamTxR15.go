package ies

import "rrc/utils"

// SL-Parameters-v1530-sl-64QAM-Tx-r15 ::= ENUMERATED
type SlParametersV1530Sl64qamTxR15 struct {
	Value utils.ENUMERATED
}

const (
	SlParametersV1530Sl64qamTxR15EnumeratedNothing = iota
	SlParametersV1530Sl64qamTxR15EnumeratedSupported
)
