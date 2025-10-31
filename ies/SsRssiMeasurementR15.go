package ies

import "rrc/utils"

// SS-RSSI-Measurement-r15 ::= SEQUENCE
type SsRssiMeasurementR15 struct {
	MeasurementslotsR15 utils.BITSTRING `lb:1,ub:80`
	EndsymbolR15        utils.INTEGER   `lb:0,ub:3`
}
