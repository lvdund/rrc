package ies

import "rrc/utils"

// UE-MeasurementsAvailable-r16 ::= SEQUENCE
// Extensible
type UeMeasurementsavailableR16 struct {
	LogmeasavailableR16          *utils.BOOLEAN
	LogmeasavailablebtR16        *utils.BOOLEAN
	LogmeasavailablewlanR16      *utils.BOOLEAN
	ConnestfailinfoavailableR16  *utils.BOOLEAN
	RlfInfoavailableR16          *utils.BOOLEAN
	SuccesshoInfoavailableR17    *utils.BOOLEAN `ext`
	SiglogmeasconfigavailableR17 *utils.BOOLEAN `ext`
}
