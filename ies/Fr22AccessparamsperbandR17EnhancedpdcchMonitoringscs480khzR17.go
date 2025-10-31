package ies

import "rrc/utils"

// FR2-2-AccessParamsPerBand-r17-enhancedPDCCH-monitoringSCS-480kHz-r17 ::= ENUMERATED
type Fr22AccessparamsperbandR17EnhancedpdcchMonitoringscs480khzR17 struct {
	Value utils.ENUMERATED
}

const (
	Fr22AccessparamsperbandR17EnhancedpdcchMonitoringscs480khzR17EnumeratedNothing = iota
	Fr22AccessparamsperbandR17EnhancedpdcchMonitoringscs480khzR17EnumeratedSupported
)
