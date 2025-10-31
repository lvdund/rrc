package ies

import "rrc/utils"

// OTDOA-PositioningCapabilities-r10-interFreqRSTD-Measurement-r10 ::= ENUMERATED
type OtdoaPositioningcapabilitiesR10InterfreqrstdMeasurementR10 struct {
	Value utils.ENUMERATED
}

const (
	OtdoaPositioningcapabilitiesR10InterfreqrstdMeasurementR10EnumeratedNothing = iota
	OtdoaPositioningcapabilitiesR10InterfreqrstdMeasurementR10EnumeratedSupported
)
