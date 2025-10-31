package ies

import "rrc/utils"

// AppLayerMeasParameters-r17-ul-MeasurementReportAppLayer-Seg-r17 ::= ENUMERATED
type ApplayermeasparametersR17UlMeasurementreportapplayerSegR17 struct {
	Value utils.ENUMERATED
}

const (
	ApplayermeasparametersR17UlMeasurementreportapplayerSegR17EnumeratedNothing = iota
	ApplayermeasparametersR17UlMeasurementreportapplayerSegR17EnumeratedSupported
)
