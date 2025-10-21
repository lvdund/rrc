package ies

import "rrc/utils"

// SCG-ConfigInfo-v1310-IEs ::= SEQUENCE
type ScgConfiginfoV1310Ies struct {
	MeasresultsstdR13               *MeasresultsstdR13
	ScelltoaddmodlistmcgExtR13      *ScelltoaddmodlistextR13
	MeasresultservcelllistscgExtR13 *MeasresultservcelllistscgExtR13
	ScelltoaddmodlistscgExtR13      *ScelltoaddmodlistscgExtR13
	ScelltoreleaselistscgExtR13     *ScelltoreleaselistextR13
	Noncriticalextension            *ScgConfiginfoV1330Ies
}
