package ies

import "rrc/utils"

// DLInformationTransfer ::= SEQUENCE
type Dlinformationtransfer struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       interface{}
}
