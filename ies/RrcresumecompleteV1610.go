package ies

import "rrc/utils"

// RRCResumeComplete-v1610-IEs ::= SEQUENCE
type RrcresumecompleteV1610 struct {
	IdlemeasavailableR16       *utils.BOOLEAN
	MeasresultidleeutraR16     *MeasresultidleeutraR16
	MeasresultidlenrR16        *MeasresultidlenrR16
	ScgResponseR16             *RrcresumecompleteV1610IesScgResponseR16
	UeMeasurementsavailableR16 *UeMeasurementsavailableR16
	MobilityhistoryavailR16    *utils.BOOLEAN
	MobilitystateR16           *RrcresumecompleteV1610IesMobilitystateR16
	NeedforgapsinfonrR16       *NeedforgapsinfonrR16
	Noncriticalextension       *RrcresumecompleteV1640
}
