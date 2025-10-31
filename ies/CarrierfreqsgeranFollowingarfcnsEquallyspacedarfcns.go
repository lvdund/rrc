package ies

import "rrc/utils"

// CarrierFreqsGERAN-followingARFCNs-equallySpacedARFCNs ::= SEQUENCE
type CarrierfreqsgeranFollowingarfcnsEquallyspacedarfcns struct {
	ArfcnSpacing            utils.INTEGER `lb:0,ub:8`
	Numberoffollowingarfcns utils.INTEGER `lb:0,ub:31`
}
