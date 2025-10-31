package ies

// FeatureSetUplink-v1630 ::= SEQUENCE
type FeaturesetuplinkV1630 struct {
	OffsetsrsCbPuschAntSwitchFr1R16                    *FeaturesetuplinkV1630OffsetsrsCbPuschAntSwitchFr1R16
	OffsetsrsCbPuschPdcchMonitorsingleoccFr1R16        *FeaturesetuplinkV1630OffsetsrsCbPuschPdcchMonitorsingleoccFr1R16
	OffsetsrsCbPuschPdcchMonitoranyoccwithoutgapFr1R16 *FeaturesetuplinkV1630OffsetsrsCbPuschPdcchMonitoranyoccwithoutgapFr1R16
	OffsetsrsCbPuschPdcchMonitoranyoccwithgapFr1R16    *FeaturesetuplinkV1630OffsetsrsCbPuschPdcchMonitoranyoccwithgapFr1R16
	Dummy                                              *FeaturesetuplinkV1630Dummy
	PartialcancellationpucchPuschPrachTxR16            *FeaturesetuplinkV1630PartialcancellationpucchPuschPrachTxR16
}
