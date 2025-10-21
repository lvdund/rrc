package ies

import "rrc/utils"

// NPUSCH-ConfigDedicated-NB-r13 ::= SEQUENCE
type NpuschConfigdedicatedNbR13 struct {
	AckNackNumrepetitionsR13 *AckNackNumrepetitionsNbR13
	NpuschAllsymbolsR13      *bool
	GrouphoppingdisabledR13  *utils.ENUMERATED
}
