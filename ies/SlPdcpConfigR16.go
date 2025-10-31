package ies

import "rrc/utils"

// SL-PDCP-Config-r16 ::= SEQUENCE
// Extensible
type SlPdcpConfigR16 struct {
	SlDiscardtimerR16    *SlPdcpConfigR16SlDiscardtimerR16
	SlPdcpSnSizeR16      *SlPdcpConfigR16SlPdcpSnSizeR16
	SlOutoforderdelivery *utils.BOOLEAN
}
