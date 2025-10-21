package ies

import "rrc/utils"

// NPUSCH-ConfigCommon-NB-r13 ::= SEQUENCE
type NpuschConfigcommonNbR13 struct {
	AckNackNumrepetitionsMsg4R13 interface{}
	SrsSubframeconfigR13         *utils.ENUMERATED
	DmrsConfigR13                *interface{}
	UlReferencesignalsnpuschR13  UlReferencesignalsnpuschNbR13
}
