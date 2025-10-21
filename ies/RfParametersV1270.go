package ies

import "rrc/utils"

// RF-Parameters-v1270 ::= SEQUENCE
type RfParametersV1270 struct {
	SupportedbandcombinationV1270    *SupportedbandcombinationV1270
	SupportedbandcombinationaddV1270 *SupportedbandcombinationaddV1270
}
