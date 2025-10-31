package ies

import "rrc/utils"

// SL-BWP-Generic-r16 ::= SEQUENCE
// Extensible
type SlBwpGenericR16 struct {
	SlBwpR16                     *Bwp
	SlLengthsymbolsR16           *SlBwpGenericR16SlLengthsymbolsR16
	SlStartsymbolR16             *SlBwpGenericR16SlStartsymbolR16
	SlPsbchConfigR16             *utils.Setuprelease[SlPsbchConfigR16]
	SlTxdirectcurrentlocationR16 *utils.INTEGER `lb:0,ub:3301`
}
