package ies

import "rrc/utils"

// AppLayerMeasParameters-r17-qoe-Streaming-MeasReport-r17 ::= ENUMERATED
type ApplayermeasparametersR17QoeStreamingMeasreportR17 struct {
	Value utils.ENUMERATED
}

const (
	ApplayermeasparametersR17QoeStreamingMeasreportR17EnumeratedNothing = iota
	ApplayermeasparametersR17QoeStreamingMeasreportR17EnumeratedSupported
)
