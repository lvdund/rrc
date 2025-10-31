package ies

// FeatureSetUplink-v1610 ::= SEQUENCE
type FeaturesetuplinkV1610 struct {
	PuschRepetitiontypebR16                      *FeaturesetuplinkV1610PuschRepetitiontypebR16
	UlCancellationselfcarrierR16                 *FeaturesetuplinkV1610UlCancellationselfcarrierR16
	UlCancellationcrosscarrierR16                *FeaturesetuplinkV1610UlCancellationcrosscarrierR16
	UlFullpwrmode2MaxsrsResinsetR16              *FeaturesetuplinkV1610UlFullpwrmode2MaxsrsResinsetR16
	CbgpuschProcessingtype1DifferenttbPerslotR16 *FeaturesetuplinkV1610CbgpuschProcessingtype1DifferenttbPerslotR16
	CbgpuschProcessingtype2DifferenttbPerslotR16 *FeaturesetuplinkV1610CbgpuschProcessingtype2DifferenttbPerslotR16
	SupportedsrsPosresourcesR16                  *SrsAllposresourcesR16
	IntrafreqdapsUlR16                           *FeaturesetuplinkV1610IntrafreqdapsUlR16
	IntrabandfreqseparationulV1620               *FreqseparationclassulV1620
	MultipucchR16                                *FeaturesetuplinkV1610MultipucchR16
	TwopucchType1R16                             *FeaturesetuplinkV1610TwopucchType1R16
	TwopucchType2R16                             *FeaturesetuplinkV1610TwopucchType2R16
	TwopucchType3R16                             *FeaturesetuplinkV1610TwopucchType3R16
	TwopucchType4R16                             *FeaturesetuplinkV1610TwopucchType4R16
	MuxSrHarqAckR16                              *FeaturesetuplinkV1610MuxSrHarqAckR16
	Dummy1                                       *FeaturesetuplinkV1610Dummy1
	Dummy2                                       *FeaturesetuplinkV1610Dummy2
	TwopucchType5R16                             *FeaturesetuplinkV1610TwopucchType5R16
	TwopucchType6R16                             *FeaturesetuplinkV1610TwopucchType6R16
	TwopucchType7R16                             *FeaturesetuplinkV1610TwopucchType7R16
	TwopucchType8R16                             *FeaturesetuplinkV1610TwopucchType8R16
	TwopucchType9R16                             *FeaturesetuplinkV1610TwopucchType9R16
	TwopucchType10R16                            *FeaturesetuplinkV1610TwopucchType10R16
	TwopucchType11R16                            *FeaturesetuplinkV1610TwopucchType11R16
	UlIntraueMuxR16                              *FeaturesetuplinkV1610UlIntraueMuxR16
	UlFullpwrmodeR16                             *FeaturesetuplinkV1610UlFullpwrmodeR16
	CrosscarrierschedulingprocessingDiffscsR16   *FeaturesetuplinkV1610CrosscarrierschedulingprocessingDiffscsR16
	UlFullpwrmode1R16                            *FeaturesetuplinkV1610UlFullpwrmode1R16
	UlFullpwrmode2SrsconfigDiffnumsrsportsR16    *FeaturesetuplinkV1610UlFullpwrmode2SrsconfigDiffnumsrsportsR16
	UlFullpwrmode2TpmigroupR16                   *FeaturesetuplinkV1610UlFullpwrmode2TpmigroupR16
}
