package ies

import "rrc/utils"

// RRCResume-v1610-IEs ::= SEQUENCE
type RrcresumeV1610 struct {
	IdlemodemeasurementreqR16 *utils.BOOLEAN
	RestoremcgScellsR16       *utils.BOOLEAN
	RestorescgR16             *utils.BOOLEAN
	MrdcSecondarycellgroupR16 *RrcresumeV1610IesMrdcSecondarycellgroupR16
	NeedforgapsconfignrR16    *utils.Setuprelease[NeedforgapsconfignrR16]
	Noncriticalextension      *RrcresumeV1700
}
