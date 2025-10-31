package ies

import "rrc/utils"

// Sensor-LocationInfo-r16 ::= SEQUENCE
// Extensible
type SensorLocationinfoR16 struct {
	SensorMeasurementinformationR16 *utils.OCTETSTRING
	SensorMotioninformationR16      *utils.OCTETSTRING
}
