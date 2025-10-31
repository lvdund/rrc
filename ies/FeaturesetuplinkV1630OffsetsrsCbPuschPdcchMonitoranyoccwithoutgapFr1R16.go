package ies

import "rrc/utils"

// FeatureSetUplink-v1630-offsetSRS-CB-PUSCH-PDCCH-MonitorAnyOccWithoutGap-fr1-r16 ::= ENUMERATED
type FeaturesetuplinkV1630OffsetsrsCbPuschPdcchMonitoranyoccwithoutgapFr1R16 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkV1630OffsetsrsCbPuschPdcchMonitoranyoccwithoutgapFr1R16EnumeratedNothing = iota
	FeaturesetuplinkV1630OffsetsrsCbPuschPdcchMonitoranyoccwithoutgapFr1R16EnumeratedSupported
)
