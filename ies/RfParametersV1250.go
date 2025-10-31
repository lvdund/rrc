package ies

// RF-Parameters-v1250 ::= SEQUENCE
type RfParametersV1250 struct {
	SupportedbandlisteutraV1250      *SupportedbandlisteutraV1250
	SupportedbandcombinationV1250    *SupportedbandcombinationV1250
	SupportedbandcombinationaddV1250 *SupportedbandcombinationaddV1250
	FreqbandpriorityadjustmentR12    *RfParametersV1250FreqbandpriorityadjustmentR12
}
