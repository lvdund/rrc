package ies

import "rrc/utils"

// RF-Parameters-v1610 ::= SEQUENCE
type RfParametersV1610 struct {
	SupportedbandcombinationV1610        *SupportedbandcombinationV1610
	SupportedbandcombinationaddV1610     *SupportedbandcombinationaddV1610
	SupportedbandcombinationreducedV1610 *SupportedbandcombinationreducedV1610
}
