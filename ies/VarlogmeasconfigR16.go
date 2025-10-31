package ies

import "rrc/utils"

// VarLogMeasConfig-r16-IEs ::= SEQUENCE
type VarlogmeasconfigR16 struct {
	AreaconfigurationR16   *AreaconfigurationR16
	BtNamelistR16          *BtNamelistR16
	WlanNamelistR16        *WlanNamelistR16
	SensorNamelistR16      *SensorNamelistR16
	LoggingdurationR16     LoggingdurationR16
	Reporttype             VarlogmeasconfigR16IesReporttype
	EarlymeasindicationR17 *utils.BOOLEAN
	AreaconfigurationV1700 *AreaconfigurationV1700
}
