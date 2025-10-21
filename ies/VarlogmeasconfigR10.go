package ies

import "rrc/utils"

// VarLogMeasConfig-r10 ::= SEQUENCE
type VarlogmeasconfigR10 struct {
	AreaconfigurationR10 *AreaconfigurationR10
	LoggingdurationR10   LoggingdurationR10
	LoggingintervalR10   LoggingintervalR10
}
