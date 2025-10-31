package ies

import "rrc/utils"

// AppLayerMeasParameters-r17-ran-VisibleQoE-Streaming-MeasReport-r17 ::= ENUMERATED
type ApplayermeasparametersR17RanVisibleqoeStreamingMeasreportR17 struct {
	Value utils.ENUMERATED
}

const (
	ApplayermeasparametersR17RanVisibleqoeStreamingMeasreportR17EnumeratedNothing = iota
	ApplayermeasparametersR17RanVisibleqoeStreamingMeasreportR17EnumeratedSupported
)
