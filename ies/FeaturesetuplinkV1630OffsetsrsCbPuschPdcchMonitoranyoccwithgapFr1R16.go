package ies

import "rrc/utils"

// FeatureSetUplink-v1630-offsetSRS-CB-PUSCH-PDCCH-MonitorAnyOccWithGap-fr1-r16 ::= ENUMERATED
type FeaturesetuplinkV1630OffsetsrsCbPuschPdcchMonitoranyoccwithgapFr1R16 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkV1630OffsetsrsCbPuschPdcchMonitoranyoccwithgapFr1R16EnumeratedNothing = iota
	FeaturesetuplinkV1630OffsetsrsCbPuschPdcchMonitoranyoccwithgapFr1R16EnumeratedSupported
)
