package ies

import "rrc/utils"

// OTDOA-PositioningCapabilities-r10 ::= SEQUENCE
type OtdoaPositioningcapabilitiesR10 struct {
	OtdoaUeAssistedR10          utils.ENUMERATED
	InterfreqrstdMeasurementR10 *utils.ENUMERATED
}
