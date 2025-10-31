package ies

import "rrc/utils"

// SL-TxProfile-r17 ::= ENUMERATED
type SlTxprofileR17 struct {
	Value utils.ENUMERATED
}

const (
	SlTxprofileR17EnumeratedNothing = iota
	SlTxprofileR17EnumeratedDrx_Compatible
	SlTxprofileR17EnumeratedDrx_Incompatible
	SlTxprofileR17EnumeratedSpare6
	SlTxprofileR17EnumeratedSpare5
	SlTxprofileR17EnumeratedSpare4
	SlTxprofileR17EnumeratedSpare3
	SlTxprofileR17EnumeratedSpare2
	SlTxprofileR17EnumeratedSpare1
)
