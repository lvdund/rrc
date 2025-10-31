package ies

import "rrc/utils"

// LoggedMeasurementConfiguration-v1700-IEs ::= SEQUENCE
type LoggedmeasurementconfigurationV1700 struct {
	SigloggedmeastypeR17   *utils.BOOLEAN
	EarlymeasindicationR17 *utils.BOOLEAN
	AreaconfigurationV1700 *AreaconfigurationV1700
	Noncriticalextension   *LoggedmeasurementconfigurationV1700IesNoncriticalextension
}
