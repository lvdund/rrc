package ies

import "rrc/utils"

// AppLayerMeasParameters-r17-qoe-VR-MeasReport-r17 ::= ENUMERATED
type ApplayermeasparametersR17QoeVrMeasreportR17 struct {
	Value utils.ENUMERATED
}

const (
	ApplayermeasparametersR17QoeVrMeasreportR17EnumeratedNothing = iota
	ApplayermeasparametersR17QoeVrMeasreportR17EnumeratedSupported
)
