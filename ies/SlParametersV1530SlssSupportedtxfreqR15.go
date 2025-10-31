package ies

import "rrc/utils"

// SL-Parameters-v1530-slss-SupportedTxFreq-r15 ::= ENUMERATED
type SlParametersV1530SlssSupportedtxfreqR15 struct {
	Value utils.ENUMERATED
}

const (
	SlParametersV1530SlssSupportedtxfreqR15EnumeratedNothing = iota
	SlParametersV1530SlssSupportedtxfreqR15EnumeratedSingle
	SlParametersV1530SlssSupportedtxfreqR15EnumeratedMultiple
)
