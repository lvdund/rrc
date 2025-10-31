package ies

// LoggedMeasurementConfiguration-v1130-IEs ::= SEQUENCE
type LoggedmeasurementconfigurationV1130 struct {
	PlmnIdentitylistR11    *PlmnIdentitylist3R11
	AreaconfigurationV1130 *AreaconfigurationV1130
	Noncriticalextension   *LoggedmeasurementconfigurationV1250
}
