package ies

import "rrc/utils"

// RRCConnectionReconfiguration-v1610-IEs ::= SEQUENCE
type RrcconnectionreconfigurationV1610 struct {
	ConditionalreconfigurationR16 *ConditionalreconfigurationR16
	DapsSourcereleaseR16          *bool
	TdmPatternconfig2R16          *TdmPatternconfigR15
	SlConfigdedicatedfornrR16     *utils.OCTETSTRING
	SlSsbPriorityeutraR16         *utils.INTEGER `lb:0,ub:8`
	Noncriticalextension          *RrcconnectionreconfigurationV1610IesNoncriticalextension
}
