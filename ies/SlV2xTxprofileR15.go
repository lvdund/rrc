package ies

import "rrc/utils"

// SL-V2X-TxProfile-r15 ::= utils.ENUMERATED // Extensible
type SlV2xTxprofileR15 struct {
	Value utils.ENUMERATED
}

const (
	SlV2xTxprofileR15EnumeratedNothing = iota
	SlV2xTxprofileR15EnumeratedRel14
	SlV2xTxprofileR15EnumeratedRel15
	SlV2xTxprofileR15EnumeratedSpare6
	SlV2xTxprofileR15EnumeratedSpare5
	SlV2xTxprofileR15EnumeratedSpare4
	SlV2xTxprofileR15EnumeratedSpare3
	SlV2xTxprofileR15EnumeratedSpare2
	SlV2xTxprofileR15EnumeratedSpare1
)
