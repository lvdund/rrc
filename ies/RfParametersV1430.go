package ies

// RF-Parameters-v1430 ::= SEQUENCE
type RfParametersV1430 struct {
	SupportedbandcombinationV1430        *SupportedbandcombinationV1430
	SupportedbandcombinationaddV1430     *SupportedbandcombinationaddV1430
	SupportedbandcombinationreducedV1430 *SupportedbandcombinationreducedV1430
	EnbRequestedparametersV1430          *RfParametersV1430EnbRequestedparametersV1430
	DifffallbackcombreportR14            *RfParametersV1430DifffallbackcombreportR14
}
