package ies

import "rrc/utils"

// SL-Parameters-v1530-sl-TxDiversity-r15 ::= ENUMERATED
type SlParametersV1530SlTxdiversityR15 struct {
	Value utils.ENUMERATED
}

const (
	SlParametersV1530SlTxdiversityR15EnumeratedNothing = iota
	SlParametersV1530SlTxdiversityR15EnumeratedSupported
)
