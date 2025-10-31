package ies

import "rrc/utils"

// SystemInformationBlockType1-v1530-IEs-crs-IntfMitigConfig-r15 ::= CHOICE
const (
	Systeminformationblocktype1V1530IesCrsIntfmitigconfigR15ChoiceNothing = iota
	Systeminformationblocktype1V1530IesCrsIntfmitigconfigR15ChoiceCrsIntfmitigenabled
	Systeminformationblocktype1V1530IesCrsIntfmitigconfigR15ChoiceCrsIntfmitignumprbs
)

type Systeminformationblocktype1V1530IesCrsIntfmitigconfigR15 struct {
	Choice              uint64
	CrsIntfmitigenabled *struct{}
	CrsIntfmitignumprbs *utils.ENUMERATED
}
