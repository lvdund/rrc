package ies

import "rrc/utils"

// MeasParameters-v1310-rssi-AndChannelOccupancyReporting-r13 ::= ENUMERATED
type MeasparametersV1310RssiAndchanneloccupancyreportingR13 struct {
	Value utils.ENUMERATED
}

const (
	MeasparametersV1310RssiAndchanneloccupancyreportingR13EnumeratedNothing = iota
	MeasparametersV1310RssiAndchanneloccupancyreportingR13EnumeratedSupported
)
