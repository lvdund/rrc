package ies

import "rrc/utils"

// FR2-2-AccessParamsPerBand-r17-enhancedPDCCH-monitoringSCS-960kHz-r17-pdcch-monitoring4-2-r17 ::= ENUMERATED
type Fr22AccessparamsperbandR17EnhancedpdcchMonitoringscs960khzR17PdcchMonitoring42R17 struct {
	Value utils.ENUMERATED
}

const (
	Fr22AccessparamsperbandR17EnhancedpdcchMonitoringscs960khzR17PdcchMonitoring42R17EnumeratedNothing = iota
	Fr22AccessparamsperbandR17EnhancedpdcchMonitoringscs960khzR17PdcchMonitoring42R17EnumeratedSupported
)
