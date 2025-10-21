package ies

import "rrc/utils"

// RF-Parameters-v1310 ::= SEQUENCE
type RfParametersV1310 struct {
	EnbRequestedparametersR13          *interface{}
	MaximumccsretrievalR13             *utils.ENUMERATED
	SkipfallbackcombinationsR13        *utils.ENUMERATED
	ReducedintnoncontcombR13           *utils.ENUMERATED
	SupportedbandlisteutraV1310        *SupportedbandlisteutraV1310
	SupportedbandcombinationreducedR13 *SupportedbandcombinationreducedR13
}
