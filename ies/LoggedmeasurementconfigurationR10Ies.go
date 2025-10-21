package ies

import "rrc/utils"

// LoggedMeasurementConfiguration-r10-IEs ::= SEQUENCE
type LoggedmeasurementconfigurationR10Ies struct {
	TracereferenceR10           TracereferenceR10
	TracerecordingsessionrefR10 utils.OCTETSTRING
	TceIdR10                    utils.OCTETSTRING
	AbsolutetimeinfoR10         AbsolutetimeinfoR10
	AreaconfigurationR10        *AreaconfigurationR10
	LoggingdurationR10          LoggingdurationR10
	LoggingintervalR10          LoggingintervalR10
	Noncriticalextension        *LoggedmeasurementconfigurationV1080Ies
}
