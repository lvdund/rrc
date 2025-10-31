package ies

import "rrc/utils"

// MeasParameters-NB-r16-dl-ChannelQualityReporting-r16 ::= ENUMERATED
type MeasparametersNbR16DlChannelqualityreportingR16 struct {
	Value utils.ENUMERATED
}

const (
	MeasparametersNbR16DlChannelqualityreportingR16EnumeratedNothing = iota
	MeasparametersNbR16DlChannelqualityreportingR16EnumeratedSupported
)
