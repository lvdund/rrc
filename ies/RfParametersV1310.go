package ies

// RF-Parameters-v1310 ::= SEQUENCE
type RfParametersV1310 struct {
	EnbRequestedparametersR13          *RfParametersV1310EnbRequestedparametersR13
	MaximumccsretrievalR13             *RfParametersV1310MaximumccsretrievalR13
	SkipfallbackcombinationsR13        *RfParametersV1310SkipfallbackcombinationsR13
	ReducedintnoncontcombR13           *RfParametersV1310ReducedintnoncontcombR13
	SupportedbandlisteutraV1310        *SupportedbandlisteutraV1310
	SupportedbandcombinationreducedR13 *SupportedbandcombinationreducedR13
}
