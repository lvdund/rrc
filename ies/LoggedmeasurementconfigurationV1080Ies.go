package ies

import "rrc/utils"

// LoggedMeasurementConfiguration-v1080-IEs ::= SEQUENCE
type LoggedmeasurementconfigurationV1080Ies struct {
	LatenoncriticalextensionR10 *utils.OCTETSTRING
	Noncriticalextension        *LoggedmeasurementconfigurationV1130Ies
}
