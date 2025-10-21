package ies

import "rrc/utils"

// RF-Parameters-v1470 ::= SEQUENCE
type RfParametersV1470 struct {
	SupportedbandcombinationV1470        *SupportedbandcombinationV1470
	SupportedbandcombinationaddV1470     *SupportedbandcombinationaddV1470
	SupportedbandcombinationreducedV1470 *SupportedbandcombinationreducedV1470
}
