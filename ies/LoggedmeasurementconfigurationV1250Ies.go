package ies

import "rrc/utils"

// LoggedMeasurementConfiguration-v1250-IEs ::= SEQUENCE
type LoggedmeasurementconfigurationV1250Ies struct {
	TargetmbsfnArealistR12 *TargetmbsfnArealistR12
	Noncriticalextension   *LoggedmeasurementconfigurationV1530Ies
}
