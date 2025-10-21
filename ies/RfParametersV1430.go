package ies

import "rrc/utils"

// RF-Parameters-v1430 ::= SEQUENCE
type RfParametersV1430 struct {
	SupportedbandcombinationV1430        *SupportedbandcombinationV1430
	SupportedbandcombinationaddV1430     *SupportedbandcombinationaddV1430
	SupportedbandcombinationreducedV1430 *SupportedbandcombinationreducedV1430
	EnbRequestedparametersV1430          *interface{}
	DifffallbackcombreportR14            *utils.ENUMERATED
}
