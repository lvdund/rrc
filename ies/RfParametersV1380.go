package ies

import "rrc/utils"

// RF-Parameters-v1380 ::= SEQUENCE
type RfParametersV1380 struct {
	SupportedbandcombinationV1380        *SupportedbandcombinationV1380
	SupportedbandcombinationaddV1380     *SupportedbandcombinationaddV1380
	SupportedbandcombinationreducedV1380 *SupportedbandcombinationreducedV1380
}
