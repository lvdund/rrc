package ies

import "rrc/utils"

// MeasParameters-v1610-ce-DL-ChannelQualityReporting-r16 ::= ENUMERATED
type MeasparametersV1610CeDlChannelqualityreportingR16 struct {
	Value utils.ENUMERATED
}

const (
	MeasparametersV1610CeDlChannelqualityreportingR16EnumeratedNothing = iota
	MeasparametersV1610CeDlChannelqualityreportingR16EnumeratedSupported
)
