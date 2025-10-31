package ies

// FeatureSetUplink-v1710 ::= SEQUENCE
type FeaturesetuplinkV1710 struct {
	MtrpPuschTypeaCbR17                *FeaturesetuplinkV1710MtrpPuschTypeaCbR17
	MtrpPuschRepetitiontypeaR17        *FeaturesetuplinkV1710MtrpPuschRepetitiontypeaR17
	MtrpPucchIntraslotR17              *FeaturesetuplinkV1710MtrpPucchIntraslotR17
	SrsAntennaswitching2sp1periodicR17 *FeaturesetuplinkV1710SrsAntennaswitching2sp1periodicR17
	SrsExtensionaperiodicsrsR17        *FeaturesetuplinkV1710SrsExtensionaperiodicsrsR17
	SrsOneapSrsR17                     *FeaturesetuplinkV1710SrsOneapSrsR17
	UePowerclassperbandperbcR17        *FeaturesetuplinkV1710UePowerclassperbandperbcR17
	TxSupportUlGapfr2R17               *FeaturesetuplinkV1710TxSupportUlGapfr2R17
}
