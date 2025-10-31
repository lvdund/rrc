package ies

import "rrc/utils"

// CarrierFreqsGERAN-followingARFCNs ::= CHOICE
const (
	CarrierfreqsgeranFollowingarfcnsChoiceNothing = iota
	CarrierfreqsgeranFollowingarfcnsChoiceExplicitlistofarfcns
	CarrierfreqsgeranFollowingarfcnsChoiceEquallyspacedarfcns
	CarrierfreqsgeranFollowingarfcnsChoiceVariablebitmapofarfcns
)

type CarrierfreqsgeranFollowingarfcns struct {
	Choice                 uint64
	Explicitlistofarfcns   *Explicitlistofarfcns
	Equallyspacedarfcns    *CarrierfreqsgeranFollowingarfcnsEquallyspacedarfcns
	Variablebitmapofarfcns *utils.OCTETSTRING `lb:1,ub:16`
}
