package ies

import "rrc/utils"

// LoggedMeasurementConfiguration-v1130-IEs ::= SEQUENCE
type LoggedmeasurementconfigurationV1130Ies struct {
	PlmnIdentitylistR11    *PlmnIdentitylist3R11
	AreaconfigurationV1130 *AreaconfigurationV1130
	Noncriticalextension   *LoggedmeasurementconfigurationV1250Ies
}
