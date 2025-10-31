package ies

import "rrc/utils"

// MeasurementTimingConfiguration-v1550-IEs ::= SEQUENCE
type MeasurementtimingconfigurationV1550 struct {
	Camponfirstssb       utils.BOOLEAN
	Pscellonlyonfirstssb utils.BOOLEAN
	Noncriticalextension *MeasurementtimingconfigurationV1610
}
