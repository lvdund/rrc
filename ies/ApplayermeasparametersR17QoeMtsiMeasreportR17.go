package ies

import "rrc/utils"

// AppLayerMeasParameters-r17-qoe-MTSI-MeasReport-r17 ::= ENUMERATED
type ApplayermeasparametersR17QoeMtsiMeasreportR17 struct {
	Value utils.ENUMERATED
}

const (
	ApplayermeasparametersR17QoeMtsiMeasreportR17EnumeratedNothing = iota
	ApplayermeasparametersR17QoeMtsiMeasreportR17EnumeratedSupported
)
