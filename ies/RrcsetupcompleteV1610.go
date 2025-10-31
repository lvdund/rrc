package ies

import "rrc/utils"

// RRCSetupComplete-v1610-IEs ::= SEQUENCE
type RrcsetupcompleteV1610 struct {
	IabNodeindicationR16       *utils.BOOLEAN
	IdlemeasavailableR16       *utils.BOOLEAN
	UeMeasurementsavailableR16 *UeMeasurementsavailableR16
	MobilityhistoryavailR16    *utils.BOOLEAN
	MobilitystateR16           *RrcsetupcompleteV1610IesMobilitystateR16
	Noncriticalextension       *RrcsetupcompleteV1690
}
