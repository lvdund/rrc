package ies

// FeatureSetDownlink ::= SEQUENCE
type Featuresetdownlink struct {
	Featuresetlistperdownlinkcc            []FeaturesetdownlinkperccId `lb:1,ub:maxNrofServingCells`
	Intrabandfreqseparationdl              *Freqseparationclass
	Scalingfactor                          *FeaturesetdownlinkScalingfactor
	Dummy8                                 *FeaturesetdownlinkDummy8
	Scellwithoutssb                        *FeaturesetdownlinkScellwithoutssb
	CsiRsMeasscellwithoutssb               *FeaturesetdownlinkCsiRsMeasscellwithoutssb
	Dummy1                                 *FeaturesetdownlinkDummy1
	Type13Css                              *FeaturesetdownlinkType13Css
	PdcchMonitoringanyoccasions            *FeaturesetdownlinkPdcchMonitoringanyoccasions
	Dummy2                                 *FeaturesetdownlinkDummy2
	UeSpecificulDlAssignment               *FeaturesetdownlinkUeSpecificulDlAssignment
	SearchspacesharingcaDl                 *FeaturesetdownlinkSearchspacesharingcaDl
	Timedurationforqcl                     *FeaturesetdownlinkTimedurationforqcl
	PdschProcessingtype1DifferenttbPerslot *FeaturesetdownlinkPdschProcessingtype1DifferenttbPerslot
	Dummy3                                 *Dummya
	Dummy4                                 *[]Dummyb `lb:1,ub:maxNrofCodebooks`
	Dummy5                                 *[]Dummyc `lb:1,ub:maxNrofCodebooks`
	Dummy6                                 *[]Dummyd `lb:1,ub:maxNrofCodebooks`
	Dummy7                                 *[]Dummye `lb:1,ub:maxNrofCodebooks`
}
