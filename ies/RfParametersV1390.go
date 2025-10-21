package ies

import "rrc/utils"

// RF-Parameters-v1390 ::= SEQUENCE
type RfParametersV1390 struct {
	SupportedbandcombinationV1390        *SupportedbandcombinationV1390
	SupportedbandcombinationaddV1390     *SupportedbandcombinationaddV1390
	SupportedbandcombinationreducedV1390 *SupportedbandcombinationreducedV1390
}
