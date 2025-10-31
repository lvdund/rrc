package ies

// FeatureSetUplink-v1640 ::= SEQUENCE
type FeaturesetuplinkV1640 struct {
	TwoharqAckCodebookType1R16                          *SubslotConfigR16
	TwoharqAckCodebookType2R16                          *SubslotConfigR16
	OffsetsrsCbPuschPdcchMonitoranyoccwithspangapFr1R16 *FeaturesetuplinkV1640OffsetsrsCbPuschPdcchMonitoranyoccwithspangapFr1R16
}
