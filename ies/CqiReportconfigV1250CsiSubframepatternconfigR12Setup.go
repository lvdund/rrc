package ies

import "rrc/utils"

// CQI-ReportConfig-v1250-csi-SubframePatternConfig-r12-setup ::= SEQUENCE
type CqiReportconfigV1250CsiSubframepatternconfigR12Setup struct {
	CsiMeassubframesetsR12 utils.BITSTRING `lb:10,ub:10`
}
