package ies

// RF-Parameters-v1180 ::= SEQUENCE
type RfParametersV1180 struct {
	FreqbandretrievalR11           *RfParametersV1180FreqbandretrievalR11
	RequestedbandsR11              *[]FreqbandindicatorR11 `lb:1,ub:maxBands`
	SupportedbandcombinationaddR11 *SupportedbandcombinationaddR11
}
