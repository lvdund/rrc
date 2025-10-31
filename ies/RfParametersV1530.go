package ies

// RF-Parameters-v1530 ::= SEQUENCE
type RfParametersV1530 struct {
	SttiSptSupportedR15                  *RfParametersV1530SttiSptSupportedR15
	SupportedbandcombinationV1530        *SupportedbandcombinationV1530
	SupportedbandcombinationaddV1530     *SupportedbandcombinationaddV1530
	SupportedbandcombinationreducedV1530 *SupportedbandcombinationreducedV1530
	Powerclass14dbmR15                   *RfParametersV1530Powerclass14dbmR15
}
