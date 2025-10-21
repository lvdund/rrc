package ies

import "rrc/utils"

// RF-Parameters-v1630 ::= SEQUENCE
type RfParametersV1630 struct {
	SupportedbandcombinationV1630        *SupportedbandcombinationV1630
	SupportedbandcombinationaddV1630     *SupportedbandcombinationaddV1630
	SupportedbandcombinationreducedV1630 *SupportedbandcombinationreducedV1630
}
