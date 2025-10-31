package ies

// SystemInformationBlockType1-NB-v1530-tdd-Parameters-r15 ::= SEQUENCE
type Systeminformationblocktype1NbV1530TddParametersR15 struct {
	TddConfigR15            TddConfigNbR15
	TddSiCarrierinfoR15     Systeminformationblocktype1NbV1530TddParametersR15TddSiCarrierinfoR15
	TddSiSubframesbitmapR15 *DlBitmapNbR13
}
