package ies

// FeatureSetUplink ::= SEQUENCE
type Featuresetuplink struct {
	Featuresetlistperuplinkcc              []FeaturesetuplinkperccId `lb:1,ub:maxNrofServingCells`
	Scalingfactor                          *FeaturesetuplinkScalingfactor
	Dummy3                                 *FeaturesetuplinkDummy3
	Intrabandfreqseparationul              *Freqseparationclass
	SearchspacesharingcaUl                 *FeaturesetuplinkSearchspacesharingcaUl
	Dummy1                                 *Dummyi
	SupportedsrsResources                  *SrsResources
	TwopucchGroup                          *FeaturesetuplinkTwopucchGroup
	Dynamicswitchsul                       *FeaturesetuplinkDynamicswitchsul
	SimultaneoustxsulNonsul                *FeaturesetuplinkSimultaneoustxsulNonsul
	PuschProcessingtype1DifferenttbPerslot *FeaturesetuplinkPuschProcessingtype1DifferenttbPerslot
	Dummy2                                 *Dummyf
}
