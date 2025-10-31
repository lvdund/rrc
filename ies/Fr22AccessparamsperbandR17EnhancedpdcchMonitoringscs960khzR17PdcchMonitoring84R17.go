package ies

import "rrc/utils"

// FR2-2-AccessParamsPerBand-r17-enhancedPDCCH-monitoringSCS-960kHz-r17-pdcch-monitoring8-4-r17 ::= ENUMERATED
type Fr22AccessparamsperbandR17EnhancedpdcchMonitoringscs960khzR17PdcchMonitoring84R17 struct {
	Value utils.ENUMERATED
}

const (
	Fr22AccessparamsperbandR17EnhancedpdcchMonitoringscs960khzR17PdcchMonitoring84R17EnumeratedNothing = iota
	Fr22AccessparamsperbandR17EnhancedpdcchMonitoringscs960khzR17PdcchMonitoring84R17EnumeratedSupported
)
