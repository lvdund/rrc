package ies

import "rrc/utils"

// RF-Parameters-v1180 ::= SEQUENCE
type RfParametersV1180 struct {
	FreqbandretrievalR11           *utils.ENUMERATED
	RequestedbandsR11              *interface{}
	SupportedbandcombinationaddR11 *SupportedbandcombinationaddR11
}
