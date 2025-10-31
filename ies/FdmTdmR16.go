package ies

import "rrc/utils"

// FDM-TDM-r16 ::= SEQUENCE
type FdmTdmR16 struct {
	RepetitionschemeR16      FdmTdmR16RepetitionschemeR16
	StartingsymboloffsetkR16 *utils.INTEGER `lb:0,ub:7`
}
