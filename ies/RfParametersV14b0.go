package ies

import "rrc/utils"

// RF-Parameters-v14b0 ::= SEQUENCE
type RfParametersV14b0 struct {
	SupportedbandcombinationV14b0        *SupportedbandcombinationV14b0
	SupportedbandcombinationaddV14b0     *SupportedbandcombinationaddV14b0
	SupportedbandcombinationreducedV14b0 *SupportedbandcombinationreducedV14b0
}
