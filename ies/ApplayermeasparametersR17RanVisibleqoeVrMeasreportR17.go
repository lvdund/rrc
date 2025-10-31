package ies

import "rrc/utils"

// AppLayerMeasParameters-r17-ran-VisibleQoE-VR-MeasReport-r17 ::= ENUMERATED
type ApplayermeasparametersR17RanVisibleqoeVrMeasreportR17 struct {
	Value utils.ENUMERATED
}

const (
	ApplayermeasparametersR17RanVisibleqoeVrMeasreportR17EnumeratedNothing = iota
	ApplayermeasparametersR17RanVisibleqoeVrMeasreportR17EnumeratedSupported
)
