package ies

// SystemInformationBlockType1-NB-v1530 ::= SEQUENCE
type Systeminformationblocktype1NbV1530 struct {
	TddParametersR15        *Systeminformationblocktype1NbV1530TddParametersR15
	SchedulinginfolistV1530 *SchedulinginfolistNbV1530
	Noncriticalextension    *Systeminformationblocktype1NbV1610
}
