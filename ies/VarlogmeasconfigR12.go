package ies

import "rrc/utils"

// VarLogMeasConfig-r12 ::= SEQUENCE
type VarlogmeasconfigR12 struct {
	AreaconfigurationR10   *AreaconfigurationR10
	AreaconfigurationV1130 *AreaconfigurationV1130
	LoggingdurationR10     LoggingdurationR10
	LoggingintervalR10     LoggingintervalR10
	TargetmbsfnArealistR12 *TargetmbsfnArealistR12
}
