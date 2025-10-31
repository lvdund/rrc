package ies

import "rrc/utils"

// SL-Parameters-v1430-v2x-TxWithShortResvInterval-r14 ::= ENUMERATED
type SlParametersV1430V2xTxwithshortresvintervalR14 struct {
	Value utils.ENUMERATED
}

const (
	SlParametersV1430V2xTxwithshortresvintervalR14EnumeratedNothing = iota
	SlParametersV1430V2xTxwithshortresvintervalR14EnumeratedSupported
)
