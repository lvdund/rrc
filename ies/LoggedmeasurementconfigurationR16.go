package ies

import "rrc/utils"

// LoggedMeasurementConfiguration-r16-IEs ::= SEQUENCE
// Extensible
type LoggedmeasurementconfigurationR16 struct {
	TracereferenceR16           TracereferenceR16
	TracerecordingsessionrefR16 utils.OCTETSTRING `lb:2,ub:2`
	TceIdR16                    utils.OCTETSTRING `lb:1,ub:1`
	AbsolutetimeinfoR16         AbsolutetimeinfoR16
	AreaconfigurationR16        *AreaconfigurationR16
	PlmnIdentitylistR16         *PlmnIdentitylist2R16
	BtNamelistR16               *utils.Setuprelease[BtNamelistR16]
	WlanNamelistR16             *utils.Setuprelease[WlanNamelistR16]
	SensorNamelistR16           *utils.Setuprelease[SensorNamelistR16]
	LoggingdurationR16          LoggingdurationR16
	Reporttype                  LoggedmeasurementconfigurationR16IesReporttype
	Latenoncriticalextension    *utils.OCTETSTRING
	Noncriticalextension        *LoggedmeasurementconfigurationV1700
}
