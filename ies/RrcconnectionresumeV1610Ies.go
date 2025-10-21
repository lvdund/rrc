package ies

import "rrc/utils"

// RRCConnectionResume-v1610-IEs ::= SEQUENCE
type RrcconnectionresumeV1610Ies struct {
	IdlemodemeasurementreqR16     *utils.ENUMERATED
	RestoremcgScellsR16           *utils.ENUMERATED
	RestorescgR16                 *utils.ENUMERATED
	ScelltoaddmodlistR16          *ScelltoaddmodlistR16
	ScelltoreleaselistR16         *ScelltoreleaselistextR13
	ScellgrouptoreleaselistR16    *ScellgrouptoreleaselistR15
	ScellgrouptoaddmodlistR16     *ScellgrouptoaddmodlistR15
	NrSecondarycellgroupconfigR16 *utils.OCTETSTRING
	PMaxeutraR16                  *PMax
	PMaxueFr1R16                  *PMax
	TdmPatternconfigR16           *TdmPatternconfigR15
	TdmPatternconfig2R16          *TdmPatternconfigR15
	Noncriticalextension          *interface{}
}
