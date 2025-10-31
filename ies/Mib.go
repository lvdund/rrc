package ies

import "rrc/utils"

// MIB ::= SEQUENCE
type Mib struct {
	Systemframenumber       utils.BITSTRING `lb:6,ub:6`
	Subcarrierspacingcommon MibSubcarrierspacingcommon
	SsbSubcarrieroffset     utils.INTEGER `lb:0,ub:15`
	DmrsTypeaPosition       MibDmrsTypeaPosition
	PdcchConfigsib1         PdcchConfigsib1
	Cellbarred              MibCellbarred
	Intrafreqreselection    MibIntrafreqreselection
	Spare                   utils.BITSTRING `lb:1,ub:1`
}
