package ies

import "rrc/utils"

// HighSpeedParameters-v1700-measurementEnhancementInterFreq-r17 ::= ENUMERATED
type HighspeedparametersV1700MeasurementenhancementinterfreqR17 struct {
	Value utils.ENUMERATED
}

const (
	HighspeedparametersV1700MeasurementenhancementinterfreqR17EnumeratedNothing = iota
	HighspeedparametersV1700MeasurementenhancementinterfreqR17EnumeratedSupported
)
