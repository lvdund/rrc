package ies

import "rrc/utils"

// RF-Parameters-v1530 ::= SEQUENCE
type RfParametersV1530 struct {
	SttiSptSupportedR15                  *utils.ENUMERATED
	SupportedbandcombinationV1530        *SupportedbandcombinationV1530
	SupportedbandcombinationaddV1530     *SupportedbandcombinationaddV1530
	SupportedbandcombinationreducedV1530 *SupportedbandcombinationreducedV1530
	Powerclass14dbmR15                   *utils.ENUMERATED
}
