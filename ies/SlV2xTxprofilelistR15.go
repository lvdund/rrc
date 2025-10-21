package ies

import "rrc/utils"

// SL-V2X-TxProfileList-r15 ::= SEQUENCE OF SL-V2X-TxProfile-r15
// SIZE (1..256)
type SlV2xTxprofilelistR15 struct {
	Value utils.Sequence[SlV2xTxprofileR15]
}
