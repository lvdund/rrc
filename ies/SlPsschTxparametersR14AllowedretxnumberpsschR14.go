package ies

import "rrc/utils"

// SL-PSSCH-TxParameters-r14-allowedRetxNumberPSSCH-r14 ::= ENUMERATED
type SlPsschTxparametersR14AllowedretxnumberpsschR14 struct {
	Value utils.ENUMERATED
}

const (
	SlPsschTxparametersR14AllowedretxnumberpsschR14EnumeratedNothing = iota
	SlPsschTxparametersR14AllowedretxnumberpsschR14EnumeratedN0
	SlPsschTxparametersR14AllowedretxnumberpsschR14EnumeratedN1
	SlPsschTxparametersR14AllowedretxnumberpsschR14EnumeratedBoth
	SlPsschTxparametersR14AllowedretxnumberpsschR14EnumeratedSpare1
)
