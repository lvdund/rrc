package ies

// NPUSCH-ConfigCommon-NB-r13 ::= SEQUENCE
type NpuschConfigcommonNbR13 struct {
	AckNackNumrepetitionsMsg4R13 []AckNackNumrepetitionsNbR13 `lb:1,ub:maxNPRACHResourcesNbR13`
	SrsSubframeconfigR13         *NpuschConfigcommonNbR13SrsSubframeconfigR13
	DmrsConfigR13                *NpuschConfigcommonNbR13DmrsConfigR13
	UlReferencesignalsnpuschR13  UlReferencesignalsnpuschNbR13
}
