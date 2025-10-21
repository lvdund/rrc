package ies

import "rrc/utils"

// SystemInformationBlockType1-NB-v1530 ::= SEQUENCE
type Systeminformationblocktype1NbV1530 struct {
	TddParametersR15        *interface{}
	SchedulinginfolistV1530 *SchedulinginfolistNbV1530
	Noncriticalextension    *Systeminformationblocktype1NbV1610
}
