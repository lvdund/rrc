package ies

// LoggedMeasurementConfiguration-v1530-IEs ::= SEQUENCE
type LoggedmeasurementconfigurationV1530 struct {
	BtNamelistR15        *BtNamelistR15
	WlanNamelistR15      *WlanNamelistR15
	Noncriticalextension *LoggedmeasurementconfigurationV1530IesNoncriticalextension
}
