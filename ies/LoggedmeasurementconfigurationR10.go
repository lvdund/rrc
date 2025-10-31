package ies

import "rrc/utils"

// LoggedMeasurementConfiguration-r10-IEs ::= SEQUENCE
type LoggedmeasurementconfigurationR10 struct {
	TracereferenceR10           TracereferenceR10
	TracerecordingsessionrefR10 utils.OCTETSTRING `lb:2,ub:2`
	TceIdR10                    utils.OCTETSTRING `lb:1,ub:1`
	AbsolutetimeinfoR10         AbsolutetimeinfoR10
	AreaconfigurationR10        *AreaconfigurationR10
	LoggingdurationR10          LoggingdurationR10
	LoggingintervalR10          LoggingintervalR10
	Noncriticalextension        *LoggedmeasurementconfigurationV1080
}
