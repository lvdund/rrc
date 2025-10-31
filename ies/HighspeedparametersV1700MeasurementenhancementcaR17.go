package ies

import "rrc/utils"

// HighSpeedParameters-v1700-measurementEnhancementCA-r17 ::= ENUMERATED
type HighspeedparametersV1700MeasurementenhancementcaR17 struct {
	Value utils.ENUMERATED
}

const (
	HighspeedparametersV1700MeasurementenhancementcaR17EnumeratedNothing = iota
	HighspeedparametersV1700MeasurementenhancementcaR17EnumeratedSupported
)
