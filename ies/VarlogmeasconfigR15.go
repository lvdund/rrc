package ies

import "rrc/utils"

// VarLogMeasConfig-r15 ::= SEQUENCE
type VarlogmeasconfigR15 struct {
	AreaconfigurationR10   *AreaconfigurationR10
	AreaconfigurationV1130 *AreaconfigurationV1130
	LoggingdurationR10     LoggingdurationR10
	LoggingintervalR10     LoggingintervalR10
	TargetmbsfnArealistR12 *TargetmbsfnArealistR12
	BtNamelistR15          *BtNamelistR15
	WlanNamelistR15        *WlanNamelistR15
}
